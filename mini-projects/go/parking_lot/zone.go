package parkinglot

import "sync"

type Zone struct {
	vehicleType VehicleType
	capacity    int
	occupied    int
	locker      *sync.Mutex
}

func NewZone(vehicleType VehicleType, capacity int) Zone {
	return Zone{
		vehicleType: vehicleType,
		capacity:    capacity,
		occupied:    0,
		locker:      &sync.Mutex{},
	}
}

func (z *Zone) Occupy() error {
	z.locker.Lock()
	defer z.locker.Unlock()

	if z.occupied >= z.capacity {
		return ErrZoneFull
	}

	z.occupied++
	return nil
}

func (z *Zone) Release() error {
	z.locker.Lock()
	defer z.locker.Unlock()

	if z.occupied == 0 {
		return ErrReleaseEmptyZone
	}

	z.occupied--
	return nil
}

func (z *Zone) CountAvailableSlots() int {
	z.locker.Lock()
	defer z.locker.Unlock()

	return z.capacity - z.occupied
}
