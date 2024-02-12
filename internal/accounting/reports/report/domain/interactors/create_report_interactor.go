package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/out"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/values"
)

type CreateReportInteractor struct {
	saveReportPort out.SaveReportOutPort
}

func NewCreateReportInteractor(
	saveReportPort out.SaveReportOutPort,
) CreateReportInteractor {
	return CreateReportInteractor{saveReportPort}
}

func (s *CreateReportInteractor) Execute(orderId uuid.UUID) (*entities.ReportEntity, error) {
	report := &entities.ReportEntity {
		ID:     uuid.New(),
		IsValid: false,
		ReportNumber: 230030,
		Positions: []entities.ReportPositionEntity{{
				ID: uuid.New(),
				Name: "empty position",
				Count: 0,
				Code: 0,
				Weight: 0,
				IsValid: false,
				Amount: values.AmountValue{
					Amount: 100,
					Rate: 0,
				},
		}},
	}


	return s.saveReportPort.Save(report)
}
