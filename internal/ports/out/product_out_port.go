package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/entities"
)
type ProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	FindByID(id uuid.UUID) (*entities.Product, error)
}
