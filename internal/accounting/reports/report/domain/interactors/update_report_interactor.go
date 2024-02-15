package interactors

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/out"
)

type UpdateReportInteractor struct {
	findReportByIdPort out.FindReportByIdOutPort
	saveReportPort out.UpdateReportOutPort
}

func NewUpdateReportInteractor(
	findReportByIdPort out.FindReportByIdOutPort,
	saveReportPort out.UpdateReportOutPort,
) UpdateReportInteractor {
	return UpdateReportInteractor{findReportByIdPort, saveReportPort}
}

func (s *UpdateReportInteractor) Execute(updateReportParams *in.UpdateReportParams) (*entities.ReportEntity, error) {
	report, err := s.findReportByIdPort.FindReportById(updateReportParams.ReportId)

	if (err != nil) {
		return nil, err
	}

	if (updateReportParams.IsValid != nil) {
		report.UpdateReportStatus(*updateReportParams.IsValid)
	}

	updatedErr := s.saveReportPort.UpdateReport(report)
	if (updatedErr != nil) {
		return nil, updatedErr
	}
	return report, nil
}
