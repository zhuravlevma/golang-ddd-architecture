package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/out"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/values"
)

type CreateReportInteractor struct {
	createReportPort out.CreateReportOutPort
}

func NewCreateReportInteractor(
	createReportPort out.CreateReportOutPort,
) CreateReportInteractor {
	return CreateReportInteractor{createReportPort}
}

func (s *CreateReportInteractor) Execute(params *in.CreateReportParams) (*entities.ReportEntity, error) {
	report := &entities.ReportEntity{
		ID:           uuid.New(),
		IsValid:      false,
		ReportNumber: 230030,
		OrderId:      params.OrderId,
		Positions: []entities.ReportPositionEntity{{
			ID:      uuid.New(),
			Name:    "empty position",
			Count:   0,
			Code:    0,
			Weight:  0,
			IsValid: false,
			Amount: values.AmountValue{
				Amount: 100,
				Rate:   0,
			},
		}},
	}

	createdErr := s.createReportPort.CreateReport(report)
	if createdErr != nil {
		return nil, createdErr
	}
	return report, nil
}
