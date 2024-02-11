package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type FindReportByIdParams struct {
	Id uuid.UUID
}

type FindReportByIdInPort interface {
	Execute(findReportByIdParams *FindReportByIdParams) (*entities.ReportEntity, error)
}
