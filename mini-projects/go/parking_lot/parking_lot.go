package parkinglot

import (
	"time"
)

type ParkingLot struct {
	zones                      map[VehicleType]*Zone
	activeTickets              map[string]Ticket // ticketID:Ticket
	activeTicketByLicencePlate map[string]string // licencePlate:ticketID
	archievedTickets           map[string]Ticket
	feeCalculator              FeeCalculator
}

func (p *ParkingLot) EnterGate(vehicle Vehicle) (Ticket, error) {
	ticketBuilder := NewTicketBuilder()

	// Step 1: validate vehicle
	if err := vehicle.Validate(); err != nil {
		return ticketBuilder.EmptyTicket(), err
	}

	// Step 2: Ensure vehicle type has parking zone
	zone, ok := p.zones[vehicle.Type()]
	if !ok {
		return ticketBuilder.EmptyTicket(), ErrInvalidVehicleType
	}

	// Step 3: reserve parking zone capacity
	if err := zone.Occupy(); err != nil {
		return ticketBuilder.EmptyTicket(), err
	}

	// Step 4: validate ticket existence
	if _, exist := p.activeTicketByLicencePlate[vehicle.LicencePlate()]; exist {
		zone.Release()
		return ticketBuilder.EmptyTicket(), ErrDuplicateTicket
	}

	// Step 5: create new ticket
	ticketBuilder.SetLicencePlate(vehicle.LicencePlate()).
		SetType(vehicle.Type()).
		SetEntryTime(time.Now())
	newTicket := ticketBuilder.Ticket()

	// Step 6: hold new ticket as active ticket
	p.activeTickets[newTicket.ID()] = newTicket
	p.activeTicketByLicencePlate[newTicket.LicencePlate()] = newTicket.ID()

	return newTicket, nil
}

func (p *ParkingLot) ExitGate(ticketID, licencePlate string) (int64, error) {
	exitTime := time.Now()

	// Step 1: search ticket details
	ticket, exist := p.activeTickets[ticketID]
	if !exist {
		return 0, ErrTicketNotFound
	}

	// Step 2: validate licen plate with ticket details
	if ticket.LicencePlate() != licencePlate {
		return 0, ErrLicencePlateMismatch
	}

	// Step 3: release parking slot
	if err := p.zones[ticket.VehicleType()].Release(); err != nil {
		return 0, err
	}

	// Step 4: remove ticket from list of active tickets
	delete(p.activeTickets, ticketID)
	delete(p.activeTicketByLicencePlate, licencePlate)

	// Step 5: calculate parking fee
	fee := p.feeCalculator.Calculate(ticket.EntryTime(), exitTime, ticket.VehicleType())

	// Step 6: update ticket details
	ticket.Close(exitTime, fee)

	// INFO: in this part we assume the system create a payment transaction.

	return fee, nil
}

func (p *ParkingLot) LostTicket(licencePlate string, vehicleType VehicleType) (int64, error) {
	/*
	 * 1. Validate from active ticket by licence plate
	 * 2. Calculate penalty fee
	 * 3. Release slot
	 */

	exitTime := time.Now()

	// Step 1: validate from active ticket
	ticketID, exist := p.activeTicketByLicencePlate[licencePlate]
	if !exist {
		return 0, ErrTicketNotFound
	}

	// Step 2: retrieve ticket details
	ticket, exist := p.activeTickets[ticketID]
	if !exist {
		return 0, ErrTicketNotFound
	}

	// Step 3: release parking slot
	if err := p.zones[vehicleType].Release(); err != nil {
		return 0, err
	}

	// Step 4: remove ticket from list of active tickets
	delete(p.activeTickets, ticketID)
	delete(p.activeTicketByLicencePlate, licencePlate)

	// Step 5: calculate penalty fee
	fee := p.feeCalculator.LostTicketFee(vehicleType)

	// Step 6: update ticket details
	ticket.Close(exitTime, fee)

	// INFO: in this part we assume the system create a payment transaction.

	return fee, nil
}

func (p *ParkingLot) CountAvailableSlots(vehicleType VehicleType) int {
	return p.zones[vehicleType].CountAvailableSlots()
}
