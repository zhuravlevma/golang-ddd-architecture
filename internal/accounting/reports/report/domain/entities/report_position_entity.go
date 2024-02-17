package entities

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/values"
)

type ReportPositionEntity struct {
	ID      uuid.UUID
	Name    string
	Count   int
	Code    int
	Weight  int
	IsValid bool
	Amount  values.AmountValue
}

func (position *ReportPositionEntity) PriceOfOnePosition() float64 {
	return position.Amount.GetAmoutWithoutTax() / float64(position.Count)
}

func (position *ReportPositionEntity) GetPriceWithTax() float64 {
	return position.PriceOfOnePosition() + position.Amount.DifferenceAfterTax()
}

func (position *ReportPositionEntity) HasNegativeDifferenceAfterTax() bool {
	return position.Amount.DifferenceAfterTax() < 0
}

func (position *ReportPositionEntity) GetValueOfTax() float64 {
	return position.Amount.DifferenceAfterTax()
}

func (position *ReportPositionEntity) UpdatePositionDiscount(discount float64) {
	position.Amount.ApplyDiscount(discount)
}

func (position *ReportPositionEntity) UpdateTaxRate(newTaxRate float64) {
	position.Amount.UpdateTaxRate(newTaxRate)
}

func (position *ReportPositionEntity) GetSumWithoutRate() float64 {
	return position.Amount.GetAmoutWithoutTax()
}

func (position *ReportPositionEntity) GetWeightOfOnePostition() int {
	return position.Weight / position.Count
}

func (position *ReportPositionEntity) HasEmptyRate() bool {
	if position.Amount.DifferenceAfterTax() == 0 {
		return true
	}
	return false
}
