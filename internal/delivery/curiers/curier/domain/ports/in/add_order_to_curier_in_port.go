package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type AddOrderToCurierParams struct {
	CurierId uuid.UUID
  OrderId uuid.UUID
}

type AddOrderToCurierInPort interface {
	Execute(addOrderToCurierParams *AddOrderToCurierParams) (*entities.CurierEntity, error)
}
