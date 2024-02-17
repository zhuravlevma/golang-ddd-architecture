package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/out"
)

type CreateOfferInteractor struct {
	createOfferOutPort out.CreateOfferOutPort
}

func (o *CreateOfferInteractor) Execute(createOfferCommand *in.CreateOfferParams) (*entities.OfferEntity, error) {
	offer := &entities.OfferEntity{
		Id:  uuid.New(),
		Name: createOfferCommand.Name,
		OrderId: createOfferCommand.OrderId,
		CurierId: nil,
		VehicleType: "Bike",
		PreferredDeliveryAreas: "New York",
		WorkingHours: "8-11",
		Weight: 0,
		Bid: 5,
	}
	err := o.createOfferOutPort.CreateOffer(offer)
	if (err != nil) {
		return nil, err
	}
	return offer, nil
}
