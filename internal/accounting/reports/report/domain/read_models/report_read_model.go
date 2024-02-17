package readmodels

import (
	"github.com/google/uuid"
)

type ReportReadModel struct {
	ID           uuid.UUID
	IsValid      bool
	OrderId      uuid.UUID
	ReportNumber int
	Positions    []ReportPositionReadModel
}

type ReportPositionReadModel struct {
	ID      uuid.UUID
	Name    string
	Count   int
	Code    int
	Weight  int
	IsValid bool
	Amount  float64
	Rate    float64
}
