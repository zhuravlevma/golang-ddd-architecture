package postgres

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/ports/out"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) out.ProductRepository {
	return &GormProductRepository{db: db}
}

func (repo *GormProductRepository) FindByID(id uuid.UUID) (*entities.Product, error) {
	var dbProduct Product
	if err := repo.db.First(&dbProduct, id).Error; err != nil {
		return nil, err
	}


	return FromDBProduct(&dbProduct), nil
}

func (repo *GormProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	dbProduct := ToDBProduct(product)

	if err := repo.db.Create(dbProduct).Error; err != nil {
		return product, err
	}

	savedEntity, err := repo.FindByID(dbProduct.ID)
	if err != nil {
		return product, err
	}


	return savedEntity, nil
}
