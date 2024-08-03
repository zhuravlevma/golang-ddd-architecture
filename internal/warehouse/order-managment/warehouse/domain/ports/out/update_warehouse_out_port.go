package out

import "github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"

type UpdateWarehouseOutPort interface {
	UpdateWarehouse(*entities.WarehouseEntity) error
}
