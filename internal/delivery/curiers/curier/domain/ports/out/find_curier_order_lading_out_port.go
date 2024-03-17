package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type FindCurierOrderLadingOutPort interface {
	FindCurierOrderLading(curierId uuid.UUID, orderId uuid.UUID) *entities.CurierEntity
}
