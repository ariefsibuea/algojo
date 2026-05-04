package parkinglot

import "strings"

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
)

func (vt *VehicleType) Validate() error {
	if *vt > Car {
		return ErrInvalidVehicleType
	}
	return nil
}

type Vehicle struct {
	licencePlate string
	_type        VehicleType
}

func NewVehicle(licencePlate string, _type VehicleType) Vehicle {
	return Vehicle{
		licencePlate: strings.TrimSpace(licencePlate),
		_type:        _type,
	}
}

func (v *Vehicle) LicencePlate() string {
	return v.licencePlate
}

func (v *Vehicle) Type() VehicleType {
	return v._type
}

func (v *Vehicle) Validate() error {
	if v == nil {
		return ErrInvalidVehicle
	}
	if v.LicencePlate() == "" {
		return ErrEmptyLicencePlate
	}

	vehicleType := v.Type()
	if err := vehicleType.Validate(); err != nil {
		return err
	}

	return nil
}
