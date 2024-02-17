package entities

import (
	"errors"

	"github.com/google/uuid"
)

type OfferEntity struct {
	Id uuid.UUID
	Name string
	OrderId uuid.UUID
	CurierId *uuid.UUID
	VehicleType string
	PreferredDeliveryAreas string
	WorkingHours string
	Weight int
	Bid int
}

func (o *OfferEntity) SetVehicleType(vehicleType string) {
	o.VehicleType = vehicleType
}

func (o *OfferEntity) SetPreferredDeliveryAreas(areas string) {
	o.PreferredDeliveryAreas = areas
}

func (o *OfferEntity) SetWorkingHours(hours string) {
	o.WorkingHours = hours
}

func (o *OfferEntity) UpdateBid() {
	if o.Weight <= 5 {
		o.Bid = 10
	} else if o.Weight <= 10 {
		o.Bid = 20
	} else {
		o.Bid = 30
	}
}

func (o *OfferEntity) IncreaseBid(amount int) {
	o.Bid += amount
}

func (o *OfferEntity) setWeight(weight int) error {
	if weight < 0 {
		return errors.New("Order weight cannot be negative")
	} else if weight <= 5 {
		o.VehicleType = "Bike"
	} else if weight <= 10 {
		o.VehicleType = "Auto"
	} else {
		o.VehicleType = "Big truck"
	}
	o.Weight = weight
	return nil
}
