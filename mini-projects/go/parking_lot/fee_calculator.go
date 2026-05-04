package parkinglot

import (
	"math"
	"time"
)

type FeeCalculator struct {
	gracePeriod    time.Duration // in minutes
	firstHour      map[VehicleType]int64
	subsequentHour map[VehicleType]int64
	dailyCap       map[VehicleType]int64
}

func (c *FeeCalculator) Calculate(entry, exit time.Time, vehicleType VehicleType) int64 {
	duration := exit.Sub(entry)
	if duration <= c.gracePeriod {
		return 0
	}

	hours := int(math.Ceil(duration.Hours()))
	if hours < 1 {
		hours = 1
	}

	fee := c.firstHour[vehicleType]
	if hours > 1 {
		fee += int64(hours-1) * c.subsequentHour[vehicleType]
	}

	if cap := c.dailyCap[vehicleType]; cap > 0 && fee > cap {
		return cap
	}

	return fee
}

func (c *FeeCalculator) LostTicketFee(vehicleType VehicleType) int64 {
	return c.dailyCap[vehicleType]
}
