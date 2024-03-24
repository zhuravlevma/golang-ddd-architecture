package orm

import (
	"github.com/google/uuid"
)

type OrderOrm struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	IsValid     bool
	WarehouseId uuid.UUID
}
