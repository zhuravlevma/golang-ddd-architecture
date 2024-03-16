package in

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type CreateOrderParams struct {
	Id string
}

type CreateOrderInPort interface {
	Execute(createOrderParams *CreateOrderParams) (*entities.CurierEntity, error)
}
