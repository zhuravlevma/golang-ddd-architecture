package orm

import (
	"github.com/google/uuid"
)

type WarehouseOrm struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	Orders []OrderOrm `gorm:"foreignKey:WarehouseId"`
}
