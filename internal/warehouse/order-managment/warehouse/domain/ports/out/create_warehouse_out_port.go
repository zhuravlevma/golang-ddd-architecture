package out

import "github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"

type CreateWarehouseOutPort interface {
	CreateWarehouse(*entities.WarehouseEntity) error
}
