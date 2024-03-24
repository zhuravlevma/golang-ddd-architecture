package in

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
)

type CreateWarehouseParams struct {
	Name string
}

type CreateWarehouseInPort interface {
	Execute(createWarehouseParams *CreateWarehouseParams) (*entities.WarehouseEntity, error)
}
