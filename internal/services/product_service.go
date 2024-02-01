package services

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/controllers/dtos"
	psql "github.com/zhuravlevma/golang-ddd-architecture/internal/db/postgres"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/entities"
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

func (s *ProductService) CreateProduct(productCommand *dtos.CreateProductDto) (*entities.Product, error) {

	var newProduct = entities.NewProduct(
		productCommand.Name,
		productCommand.Price,
	)

	err := s.productRepository.Create(newProduct)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
