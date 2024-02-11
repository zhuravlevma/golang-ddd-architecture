package orm

import (
	"github.com/google/uuid"
)

type ReportPositionOrm struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Count    int
	Code int
	Weight int
	ReportId uuid.UUID
	Sum float64
	Rate float64
	IsValid bool
}
