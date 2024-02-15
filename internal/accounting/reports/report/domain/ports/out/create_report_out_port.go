package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type CreateReportOutPort interface {
	CreateReport(reportId *entities.ReportEntity) error
}
