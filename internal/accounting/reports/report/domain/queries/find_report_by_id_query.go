package queries

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/out"
)

type FindReportByIdQuery struct {
	findReportByIdOutPort out.FindReportByIdOutPort
}

func NewFindReportByIdQuery(
	findReportByIdOutPort out.FindReportByIdOutPort,
) FindReportByIdQuery {
	return FindReportByIdQuery{findReportByIdOutPort}
}

func (s *FindReportByIdQuery) Execute(reportId uuid.UUID) (*entities.ReportEntity, error) {
	report, err := s.findReportByIdOutPort.FindReportById(reportId)
	if (err != nil) {
		return nil, err
	}

	return report, nil
}
