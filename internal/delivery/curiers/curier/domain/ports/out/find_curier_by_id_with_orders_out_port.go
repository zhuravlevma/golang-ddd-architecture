package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type FindCurierByIdWithOrdersOutPort interface {
	FindCurierByIdWithOrders(curierId uuid.UUID) entities.CurierEntity
}
