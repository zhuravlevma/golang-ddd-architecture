package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

type FindOfferByIdOutPort interface {
	FindOfferById(offerId uuid.UUID) (*entities.OfferEntity, error)
}
