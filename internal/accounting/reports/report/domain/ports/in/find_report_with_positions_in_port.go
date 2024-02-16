package in

import (
	"github.com/google/uuid"
	readmodels "github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/read_models"
)

type FindReportWithPositionsParams struct {
	Id uuid.UUID
}

type FindReportWithPositionsByIdInPort interface {
	Execute(findReportWithPositionsParams *FindReportWithPositionsParams) (*readmodels.ReportReadModel, error)
}
