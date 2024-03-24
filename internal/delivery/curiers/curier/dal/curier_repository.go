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
	var dbCurier orm.CurierOrm
	if err := repo.Db.First(&dbCurier, id).Error; err != nil {
		return nil, err
	}

	return CurierToDomain(&dbCurier), nil
}

func (repo *CurierRepository) UpdateReport(curier *entities.CurierEntity) error {
	dbCurier := CurierToOrm(curier)
	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&dbCurier).Where("id = ?", dbCurier.ID).Updates(dbCurier).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedCurier, err := repo.FindCurierById(dbCurier.ID)
	if err != nil {
		return err
	}
	*curier = *storedCurier
	return nil
}

func (repo *CurierRepository) CreateCurier(report *entities.CurierEntity) error {
	dbCurier := CurierToOrm(report)

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		if err := repo.Db.Create(dbCurier).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedCurier, err := repo.FindCurierById(dbCurier.ID)
	if err != nil {
		return err
	}
	*report = *storedCurier
	return nil
}
