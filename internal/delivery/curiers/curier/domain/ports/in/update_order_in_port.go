package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type UpdateOrderParams struct {
	CurierId    uuid.UUID
	OrderId     uuid.UUID
	Description *string
	Delivered   *bool
	Returned    *bool
}

type UpdateOrderInPort interface {
	Execute(updateOrderStatusParams *UpdateOrderParams) (*entities.CurierEntity, error)
}
