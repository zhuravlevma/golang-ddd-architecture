package entities

import (
	"errors"

	"github.com/google/uuid"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/events"
)

type OfferEntity struct {
	Id                     uuid.UUID
	Name                   string
	OrderId                uuid.UUID
	CurierId               *uuid.UUID
	VehicleType            string
	PreferredDeliveryAreas string
	WorkingHours           string
	Weight                 int
	Bid                    int
	DomainMessages         []lib.DomainMessage[lib.DomainMessagePayload]
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

func (o *OfferEntity) SetWeight(weight int) error {
	if weight < 0 {
		return errors.New("order weight cannot be negative")
	} else if weight <= 5 {
		o.VehicleType = "Bike"
	} else if weight <= 10 {
		o.VehicleType = "Auto"
	} else {
		o.VehicleType = "Big truck"
	}
	o.Weight = weight
	o.UpdateBid()
	return nil
}

func (o *OfferEntity) CurierTakeOffer(curierId uuid.UUID) {
	o.CurierId = &curierId
	o.DomainMessages = append(o.DomainMessages, events.NewOfferTakedEvent(events.OfferTakedPayload{
		OrderId: o.OrderId,
	}, o.Id))
}
