package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

type UpdateOfferOutPort interface {
	UpdateOffer(*entities.OfferEntity) error
}
