package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

type UpdateOfferParams struct {
	OfferId uuid.UUID
	CurierId *uuid.UUID
}

type UpdateOfferInPort interface {
	Execute(updateOfferCommand *UpdateOfferParams) (*entities.OfferEntity, error)
}
