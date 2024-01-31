package services

import (
	"github.com/google/uuid"
	psql "github.com/zhuravlevma/golang-ddd-architecture/internal/db/postgres"
)

type ProductService struct {
	productRepository psql.GormProductRepository
}


func NewProductService(
	productRepository psql.GormProductRepository,
) ProductService {
	return ProductService{productRepository: productRepository}
}


func (s *ProductService) FindProductByID(id uuid.UUID) (*psql.Product, error) {
	return s.productRepository.FindByID(id)
}
