package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type SaveReportOutPort interface {
	Save(reportId *entities.ReportEntity) (*entities.ReportEntity, error)
}
