package parkinglot

import "errors"

var (
	ErrZoneFull             = errors.New("zone full")
	ErrReleaseEmptyZone     = errors.New("release on empty zone")
	ErrTicketNotFound       = errors.New("ticket not found")
	ErrDuplicateTicket      = errors.New("ticket has been created")
	ErrLicencePlateMismatch = errors.New("licence plate does not match ticket")
	ErrInvalidVehicle       = errors.New("invalid vehicle")
	ErrInvalidVehicleType   = errors.New("invalid vehicle type")
	ErrEmptyLicencePlate    = errors.New("empty licence plate")
	ErrNoMatchingPlate      = errors.New("no active ticket for this plate today")
)
