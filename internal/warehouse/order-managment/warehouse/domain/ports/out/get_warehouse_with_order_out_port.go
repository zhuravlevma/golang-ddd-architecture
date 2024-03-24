package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
)

type GetWarehouseWithOrderOutPort interface {
	GetWarehouseWithOrder(warehouseId uuid.UUID, orderId uuid.UUID) entities.WarehouseEntity
}
