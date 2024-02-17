package orm

import (
	"github.com/google/uuid"
)

type ReportOrm struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	IsValid      bool
	OrderId      uuid.UUID
	ReportNumber int
	Positions    []ReportPositionOrm `gorm:"foreignKey:ReportId"`
}
