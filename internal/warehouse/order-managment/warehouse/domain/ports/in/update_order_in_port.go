package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
)

type UpdateOrderParams struct {
	WarehouseId uuid.UUID
	OrderId     uuid.UUID
	IsValid     bool
}

type UpdateOrderInPort interface {
	Execute(updateOrderParams *UpdateOrderParams) (*entities.WarehouseEntity, error)
}
