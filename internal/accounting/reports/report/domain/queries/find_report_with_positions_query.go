package queries

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/out"
	readmodels "github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/read_models"
)

type FindreportWithPositionsQuery struct {
	findReportWithPositionsByOutPort out.FindReportWithPositionsByIdOutPort
}

func NewFindreportWithPositionsQuery(
	findReportWithPositionsByOutPort out.FindReportWithPositionsByIdOutPort,
) FindreportWithPositionsQuery {
	return FindreportWithPositionsQuery{findReportWithPositionsByOutPort}
}

func (s *FindreportWithPositionsQuery) Execute(reportId uuid.UUID) (*readmodels.ReportReadModel, error) {
	reports, err := s.findReportWithPositionsByOutPort.FindReportWithPositionsId(&out.FindReportWithPositionsParams{Id: reportId})
	if (err != nil) {
		return nil, err
	}

	return reports, nil
}
