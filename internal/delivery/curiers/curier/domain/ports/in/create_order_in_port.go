package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type CreateOrderParams struct {
	Id uuid.UUID
}

type CreateOrderInPort interface {
	Execute(createOrderParams *CreateOrderParams) (*entities.CurierEntity, error)
}
