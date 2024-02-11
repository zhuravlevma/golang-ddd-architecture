package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type CreateReportParams struct {
	OrderId uuid.UUID
}

type CreateReportInPort interface {
	Execute(createReportParams *CreateReportParams) (*entities.ReportEntity, error)
}
