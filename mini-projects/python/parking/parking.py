from datetime import date, datetime

from .parking_fee_rule import ParkingFeeRule
from .parking_slot import ParkingSlot
from .report import Report
from .ticket import Ticket, TicketStatus
from .vehicle import VehicleType


class Parking:
    _default_fee_rules: dict[VehicleType, ParkingFeeRule] = {
        VehicleType.MOTORCYCLE: ParkingFeeRule(VehicleType.MOTORCYCLE, 2000),
        VehicleType.CAR: ParkingFeeRule(VehicleType.CAR, 5000),
        VehicleType.TRUCK: ParkingFeeRule(VehicleType.TRUCK, 8000),
        VehicleType.BUS: ParkingFeeRule(VehicleType.BUS, 10000),
    }

    _default_slots: dict[VehicleType, ParkingSlot] = {
        VehicleType.MOTORCYCLE: ParkingSlot(VehicleType.MOTORCYCLE, 20),
        VehicleType.CAR: ParkingSlot(VehicleType.CAR, 15),
        VehicleType.TRUCK: ParkingSlot(VehicleType.TRUCK, 10),
        VehicleType.BUS: ParkingSlot(VehicleType.BUS, 5),
    }

    def __init__(
        self,
        fee_rules: dict[VehicleType, ParkingFeeRule] | None = None,
        slots: dict[VehicleType, ParkingSlot] | None = None,
    ) -> None:

        self.tickets: dict[str, Ticket] = {}
        self.fee_rules: dict[VehicleType, ParkingFeeRule] = fee_rules if fee_rules else self._default_fee_rules
        self.slots: dict[VehicleType, ParkingSlot] = slots if slots else self._default_slots

    def tap_in(self, license_no: str, vehicle_type: str) -> Ticket | None:
        time_now = datetime.now()
        vehicle_type_ = VehicleType[vehicle_type]

        # check parking slot availability
        parking_slot = self.slots.get(vehicle_type_, None)
        if not parking_slot or parking_slot.slot == 0:
            print(f"parking slot for {vehicle_type} is not available")
            return None

        parking_fee_rule = self.fee_rules.get(vehicle_type_, None)
        if not parking_fee_rule:
            print(f"parking fee rule for {vehicle_type} is not found")
            return None

        # calculate parking fee
        parking_fee = parking_fee_rule.calculate_fee()

        # create new ticket
        new_ticket = Ticket(
            license_no=license_no,
            vehicle_type=vehicle_type_,
            entry_time=time_now,
            fee=parking_fee,
        )

        # store ticket
        self.tickets[new_ticket.ticket_no] = new_ticket
        # decrease parking slot
        self.slots[parking_slot.vehicle_type].decrease_slot()

        return new_ticket

    def tap_out(self, ticket_no: str, total_payment: int) -> Ticket | None:
        time_now = datetime.now()

        # find ticket
        ticket = self.tickets.get(ticket_no, None)
        if not ticket:
            print(f"ticket {ticket_no} is not found")
            return None

        # check payment
        total_change = total_payment - ticket.payment.total
        if total_change < 0:
            print("payment failed, total payment is insufficient")
            return None

        # update status payment
        ticket.payment.update_status_paid()
        # update status ticket
        ticket.exit_time = time_now
        ticket.update_status_paid()
        # increase parking slot
        parking_slot = self.slots.get(ticket.vehicle.type_, None)
        if parking_slot:
            self.slots[parking_slot.vehicle_type].increase_slot()
        # update ticket in data storage
        self.tickets[ticket_no] = ticket

        print(f"ticket payment success, total change: {total_change}")
        return ticket

    def fetch_tickets_by_date(self, parking_date: date) -> list[Ticket] | None:
        start_time = datetime(parking_date.year, parking_date.month, parking_date.day)
        end_time = datetime(parking_date.year, parking_date.month, parking_date.day, 23, 59, 59)

        tickets: list[Ticket] = []
        for ticket in self.tickets.values():
            if (
                ticket.exit_time
                and ticket.exit_time >= start_time
                and ticket.exit_time <= end_time
                and ticket.status == TicketStatus.PAID
            ):
                # add total payment
                tickets.append(ticket)

        return tickets

    def generate_report(self, parking_date: date) -> Report | None:
        # fetch all tickets with exit time within 'parking_date'
        tickets = self.fetch_tickets_by_date(parking_date)

        # calculate total income
        total_income = 0
        if tickets:
            for _, ticket in enumerate(tickets):
                total_income += ticket.payment.total

        return Report(parking_date, total_income)
