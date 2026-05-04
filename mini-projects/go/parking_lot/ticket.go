package parkinglot

import (
	"time"

	"github.com/google/uuid"
)

type TicketStatus int

const (
	TicketStatusEnter TicketStatus = iota
	TicketStatusOnPayment
	TicketStatusExit
)

type Ticket struct {
	id           string
	licencePlate string
	vehicleType  VehicleType
	entryTime    time.Time
	exitTime     time.Time
	fee          int64
	status       TicketStatus
}

func (t *Ticket) ID() string {
	return t.id
}

func (t *Ticket) LicencePlate() string {
	return t.licencePlate
}

func (t *Ticket) VehicleType() VehicleType {
	return t.vehicleType
}

func (t *Ticket) EntryTime() time.Time {
	return t.entryTime
}

func (t *Ticket) Close(exit time.Time, fee int64) {
	t.exitTime = exit
	t.fee = fee
	t.status = TicketStatusExit
}

type TicketBuilder struct {
	ticket *Ticket
}

func NewTicketBuilder() TicketBuilder {
	return TicketBuilder{
		ticket: &Ticket{
			id:     uuid.NewString(),
			status: TicketStatusEnter,
		},
	}
}

func (b *TicketBuilder) SetLicencePlate(licencePlate string) *TicketBuilder {
	b.ticket.licencePlate = licencePlate
	return b
}

func (b *TicketBuilder) SetType(vehicleType VehicleType) *TicketBuilder {
	b.ticket.vehicleType = vehicleType
	return b
}

func (b *TicketBuilder) SetEntryTime(entry time.Time) *TicketBuilder {
	b.ticket.entryTime = entry
	return b
}

func (b *TicketBuilder) EmptyTicket() Ticket {
	return Ticket{}
}

func (b *TicketBuilder) Ticket() Ticket {
	return *b.ticket
}
