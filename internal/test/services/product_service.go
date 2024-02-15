package services

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/entities/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/ports/out"
)

type ProductService struct {
	productRepository out.ProductRepository
}


func NewProductService(
	productRepository out.ProductRepository,
) ProductService {
	return ProductService{productRepository: productRepository}
}


func (s *ProductService) FindProductByID(id uuid.UUID) (*entities.Product, error) {
	return s.productRepository.FindByID(id)
}

func (s *ProductService) CreateProduct(productCommand *dtos.CreateProductDto) (*entities.Product, error) {

	var newProduct = entities.NewProduct(
		productCommand.Name,
		productCommand.Price,
	)

	savedProduct, err := s.productRepository.Create(newProduct)
	if err != nil {
		return nil, err
	}
	return savedProduct, nil
}
