package postgres

import "github.com/zhuravlevma/golang-ddd-architecture/internal/entities"

func ToDBProduct(product *entities.Product) *Product {
	var p = &Product{
		Name:     product.Name,
		Price:    product.Price,
	}
	p.ID = product.ID

	return p
}
