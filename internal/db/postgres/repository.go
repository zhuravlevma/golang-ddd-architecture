package postgres

import (
	"github.com/google/uuid"
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
	if err := repo.db.Preload("Seller").First(&dbProduct, id).Error; err != nil {
		return nil, err
	}

	// Map back to domain entity
	return &dbProduct, nil;
}
