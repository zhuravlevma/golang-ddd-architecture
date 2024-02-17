package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type FindReportByIdOutPort interface {
	FindReportById(reportId uuid.UUID) (*entities.ReportEntity, error)
}
