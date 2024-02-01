package postgres

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/entities"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) GormProductRepository {
	return GormProductRepository{db: db}
}

func (repo *GormProductRepository) FindByID(id uuid.UUID) (*Product, error) {
	var dbProduct Product
	if err := repo.db.First(&dbProduct, id).Error; err != nil {
		return nil, err
	}

	// Map back to domain entity
	return &dbProduct, nil;
}

func (repo *GormProductRepository) Create(product *entities.Product) error {
	// Map domain entity to DB model
	dbProduct := ToDBProduct(product)

	if err := repo.db.Create(dbProduct).Error; err != nil {
		return err
	}

	// Read row from DB to never return different data than persisted
	_, err := repo.FindByID(dbProduct.ID)
	if err != nil {
		return err
	}


	return nil
}
