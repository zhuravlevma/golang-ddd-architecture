package dal

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
	"gorm.io/gorm"
)

type CurierRepository struct {
	Db *gorm.DB
}

func (repo *CurierRepository) FindCurierById(id uuid.UUID) (*entities.CurierEntity, error) {
	var dbReport orm.CurierOrm
	if err := repo.Db.First(&dbReport, id).Error; err != nil {
		return nil, err
	}

	return CurierToDomain(&dbReport), nil
}

func (repo *CurierRepository) UpdateReport(report *entities.CurierEntity) error {
	dbReport := CurierToOrm(report)
	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entities.CurierEntity{}).Where("id = ?", dbReport.ID).Updates(dbReport).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedReport, err := repo.FindCurierById(dbReport.ID)
	if err != nil {
		return err
	}
	*report = *storedReport
	return nil
}

func (repo *CurierRepository) CreateReport(report *entities.CurierEntity) error {
	dbReport := CurierToOrm(report)

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		if err := repo.Db.Create(dbReport).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedReport, err := repo.FindCurierById(dbReport.ID)
	if err != nil {
		return err
	}
	*report = *storedReport
	return nil
}
