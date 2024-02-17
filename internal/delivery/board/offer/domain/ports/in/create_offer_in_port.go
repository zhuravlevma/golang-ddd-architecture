package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

type CreateOfferParams struct {
	Name    string
	OrderId uuid.UUID
}

type CreateOfferInPort interface {
	Execute(createOfferCommand *CreateOfferParams) (*entities.OfferEntity, error)
}
