from datetime import datetime
from enum import Enum

from .ticket_payment import TicketPayment
from .vehicle import Vehicle, VehicleType


class TicketStatus(Enum):
    ISSUED = 1
    PAID = 1


class Ticket:
    def __init__(self, license_no: str, vehicle_type: VehicleType, entry_time: datetime, fee: int) -> None:
        time_now = datetime.now()
        ticket_no = f"TKT-{int(time_now.timestamp())}"

        self.ticket_no: str = ticket_no
        self.entry_time: datetime = entry_time
        self.exit_time: datetime | None = None
        self.vehicle: Vehicle = Vehicle(vehicle_type, license_no)
        self.payment: TicketPayment = TicketPayment(ticket_no, fee)
        self.status: TicketStatus = TicketStatus.ISSUED

    def update_status_paid(self):
        self.status = TicketStatus.PAID
