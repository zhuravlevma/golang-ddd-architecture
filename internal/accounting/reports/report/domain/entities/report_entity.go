package entities

import (
	"github.com/google/uuid"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/events"
)

type ReportEntity struct {
	ID             uuid.UUID
	IsValid        bool
	OrderId        uuid.UUID
	ReportNumber   int
	Positions      []ReportPositionEntity
	DomainMessages []lib.DomainMessage
}

func (report *ReportEntity) UpdateReportStatus(status bool) {
	if status == true {
		report.IsValid = true
		report.DomainMessages = append(report.DomainMessages, events.NewReportValidatedEvent(events.ReportValidatedPayload{
			OrderId: report.OrderId,
		}, report.ID))
	} else {
		report.IsValid = false
	}
}

func (report *ReportEntity) ApplyReport() {
	if report.GetTotalAmountWithTax() > 10000 {
		report.UpdateReportStatus(true)
	}

	for _, position := range report.Positions {
		position.UpdatePositionDiscount(0.1)
	}
}

func (report *ReportEntity) GetTaxAmount() (totalTax float64) {
	for _, position := range report.Positions {
		totalTax += position.GetValueOfTax()
	}
	return totalTax
}

func (report *ReportEntity) GetPositionsAboveTaxThreshold(threshold float64) (response []*ReportPositionEntity) {
	for _, position := range report.Positions {
		if position.GetValueOfTax() > threshold {
			response = append(response, &position)
		}
	}
	return
}

func (report *ReportEntity) GetTotalAmountWithTax() (totalAmount float64) {
	for _, position := range report.Positions {
		totalAmount += position.GetPriceWithTax()
	}
	return totalAmount
}
