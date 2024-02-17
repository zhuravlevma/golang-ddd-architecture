package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

type CreateOfferOutPort interface {
	CreateOffer(*entities.OfferEntity) error
}
