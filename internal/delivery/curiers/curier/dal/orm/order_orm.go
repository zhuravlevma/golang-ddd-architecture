package orm

import (
	"github.com/google/uuid"
)

type OrderOrm struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	OrderId     uuid.UUID
	IsActive    bool
	CurierId    uuid.UUID
	TotalSum    int
	Weight      int
}
