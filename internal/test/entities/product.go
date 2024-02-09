package entities

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID
	Name   string
	Price  float64
}

func (p *Product) validate() error {
	if p.Name == "" || p.Price <= 0 {
		return errors.New("invalid product details")
	}

	return nil
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:     uuid.New(),
		Name:   name,
		Price:  price,
	}
}
