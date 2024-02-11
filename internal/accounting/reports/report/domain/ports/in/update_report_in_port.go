package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type UpdateReportParams struct {
	ReportId uuid.UUID
	IsValid *bool
}

type UpdateReportInPort interface {
	Execute(findReportWithPositionsParams *UpdateReportParams) (*entities.ReportEntity, error)
}
