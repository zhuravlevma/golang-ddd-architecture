package out

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)

type FindReportWithPositionsParams struct {
	id string
}

type FindReportWithPositionsByOutPort interface {
	findReportWithPositionsByOutInPort(findReportByIdQuery *FindReportWithPositionsParams) (*entities.ReportEntity, error)
}
