package dal

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/values"
)
func ReportToOrm(report *entities.ReportEntity) *orm.ReportOrm {
	var positions []orm.ReportPositionOrm;
	for _, entityPosition := range report.Positions {
		positions = append(positions, orm.ReportPositionOrm{
			ID: entityPosition.ID,
			Name: entityPosition.Name,
			Count: entityPosition.Count,
			Code: entityPosition.Code,
			Weight: entityPosition.Weight,
			ReportId: report.ID,
			Sum: entityPosition.Amount.Amount,
			Rate: entityPosition.Amount.Rate,
		})
	}
	return &orm.ReportOrm{
		ID:     report.ID,
		IsValid: report.IsValid,
		ReportNumber: report.ReportNumber,
		Positions: positions,
	}
}

func ReportToDomain(ormReport *orm.ReportOrm) *entities.ReportEntity {
	var positions []entities.ReportPositionEntity;
	for _, ormPosition := range ormReport.Positions {
		positions = append(positions, entities.ReportPositionEntity{
			ID: ormPosition.ID,
			Name: ormPosition.Name,
			Count: ormPosition.Count,
			Code: ormPosition.Code,
			Weight: ormPosition.Weight,
			Amount: values.AmountValue{Amount: ormPosition.Sum, Rate: ormPosition.Rate},
		})
	}
	return &entities.ReportEntity{
		ID:     ormReport.ID,
		IsValid: ormReport.IsValid,
		ReportNumber: ormReport.ReportNumber,
		Positions: positions,
	}
}
