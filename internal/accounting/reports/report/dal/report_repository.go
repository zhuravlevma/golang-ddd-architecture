package dal

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"gorm.io/gorm"
)


type ReportRepository struct {
	Db *gorm.DB
}

func (repo *ReportRepository) FindReportById(id uuid.UUID) (*entities.ReportEntity, error) {
	var dbReport orm.ReportOrm
	if err := repo.Db.First(&dbReport, id).Error; err != nil {
		return nil, err
	}


	return ReportToDomain(&dbReport), nil
}

func (repo *ReportRepository) UpdateReport(report *entities.ReportEntity) error {
	dbReport := ReportToOrm(report)
	err := repo.Db.Model(&entities.ReportEntity{}).Where("id = ?", dbReport.ID).Updates(dbReport).Error
	if err != nil {
		return err
	}
	storedReport, err := repo.FindReportById(dbReport.ID)
	if err != nil {
		return err
	}
	*report = *storedReport
	return nil
}

func (repo *ReportRepository) CreateReport(report *entities.ReportEntity) error {
	dbReport := ReportToOrm(report)
	if err := repo.Db.Create(dbReport).Error; err != nil {
		return err
	}
	storedReport, err := repo.FindReportById(dbReport.ID)
	if err != nil {
		return err
	}
	*report = *storedReport
	return nil
}

func (repo *ReportRepository) FindReportWithPositionsByOutInPort(id uuid.UUID) (reports *[]entities.ReportEntity) {
	repo.Db.Model(&entities.ReportEntity{}).Preload("Students", func(db *gorm.DB) *gorm.DB {
    return db.Where("reports.id = ?", id)
	}).Find(&reports)
	return
}
