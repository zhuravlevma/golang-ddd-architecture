package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type UpdateReportOutPort interface {
	UpdateReport(reportId *entities.ReportEntity) error
}
