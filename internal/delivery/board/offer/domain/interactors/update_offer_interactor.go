package interactors

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/out"
)

type UpdateOfferInteractor struct {
	findOfferByIdOutPort out.FindOfferByIdOutPort
	updateOfferOutPort   out.UpdateOfferOutPort
}

func (o *UpdateOfferInteractor) Execute(updateOfferCommand *in.UpdateOfferParams) (*entities.OfferEntity, error) {
	offer, err := o.findOfferByIdOutPort.FindOfferById(updateOfferCommand.OfferId)

	if err != nil {
		return nil, err
	}

	if updateOfferCommand.CurierId != nil {
		offer.CurierTakeOffer(*updateOfferCommand.CurierId)
	}

	updatedErr := o.updateOfferOutPort.UpdateOffer(offer)
	if updatedErr != nil {
		return nil, updatedErr
	}
	return offer, nil
}
