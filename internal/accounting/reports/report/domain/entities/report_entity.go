package entities

import "github.com/google/uuid"

type ReportEntity struct {
	ID     uuid.UUID
	IsValid  bool
	OrderId uuid.UUID
	ReportNumber int
	positions []ReportPositionEntity
}

func (report *ReportEntity) UpdateReportStatus(status bool) {
	if status == true {
		report.IsValid = true
		// add message to events
	} else {
		report.IsValid = false
	}
}

func (report *ReportEntity) ApplyReport() {
	if report.GetTotalAmountWithTax() > 10000 {
		report.UpdateReportStatus(true)
	}

	for _, position := range report.positions {
		position.UpdatePositionDiscount(0.1);
	}
}

func (report *ReportEntity) GetTaxAmount() (totalTax float64) {
	for _, position := range report.positions {
		totalTax += position.GetValueOfTax();
	}
	return totalTax
}

func (report *ReportEntity) GetPositionsAboveTaxThreshold(threshold float64) (response []*ReportPositionEntity) {
	for _, position := range report.positions {
		if position.GetValueOfTax() > threshold {
			response = append(response, &position)
		}
	}
	return
}

func (report *ReportEntity) GetTotalAmountWithTax() (totalAmount float64) {
	for _, position := range report.positions {
		totalAmount += position.GetPriceWithTax();
	}
	return totalAmount
}
