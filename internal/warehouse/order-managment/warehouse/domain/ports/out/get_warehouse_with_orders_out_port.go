package out

import (
	"github.com/google/uuid"
)

type GetWarehouseWithOrdersOutPort interface {
	GetWarehouseWithOrders(warehouseId uuid.UUID) error
}
