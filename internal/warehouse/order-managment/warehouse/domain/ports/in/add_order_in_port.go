package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
)

type AddOrderParams struct {
	WarehouseId uuid.UUID
	Name        string
	OrderId     uuid.UUID
	IsValid     bool
}

type AddOrderInPort interface {
	Execute(addOrderParams *AddOrderParams) (*entities.WarehouseEntity, error)
}
