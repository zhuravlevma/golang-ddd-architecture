package out

import (
	"github.com/google/uuid"
	readmodels "github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/read_models"
)

type FindReportWithPositionsParams struct {
	Id uuid.UUID
}

type FindReportWithPositionsByIdOutPort interface {
	FindReportWithPositionsId(findReportByIdQuery *FindReportWithPositionsParams) (*readmodels.ReportReadModel, error)
}
