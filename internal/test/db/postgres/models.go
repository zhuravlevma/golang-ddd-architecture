package postgres

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Price    float64
}
