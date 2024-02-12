package dal

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
	"gorm.io/gorm"
)


type GormReportRepository struct {
	db *gorm.DB
}

func (repo *GormReportRepository) FindReportById(id uuid.UUID) (*entities.ReportEntity, error) {
	var dbReport orm.ReportOrm
	if err := repo.db.First(&dbReport, id).Error; err != nil {
		return nil, err
	}


	return ReportToDomain(&dbReport), nil
}

func (repo *GormReportRepository) UpdateReport(report *entities.ReportEntity) error {
	dbReport := ReportToOrm(report)
	err := repo.db.Model(&entities.ReportEntity{}).Where("id = ?", dbReport.ID).Updates(dbReport).Error
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

func (repo *GormReportRepository) CreateReport(report *entities.ReportEntity) error {
	dbReport := ReportToOrm(report)
	if err := repo.db.Create(dbReport).Error; err != nil {
		return err
	}
	storedReport, err := repo.FindReportById(dbReport.ID)
	if err != nil {
		return err
	}
	*report = *storedReport
	return nil
}

func (repo *GormReportRepository) FindReportWithPositionsByOutInPort(id uuid.UUID) (reports *[]entities.ReportEntity) {
	repo.db.Model(&entities.ReportEntity{}).Preload("Students", func(db *gorm.DB) *gorm.DB {
    return db.Where("reports.id = ?", id)
	}).Find(&reports)
	return
}
