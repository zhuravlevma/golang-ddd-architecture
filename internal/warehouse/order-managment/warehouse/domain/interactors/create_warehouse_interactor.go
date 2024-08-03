package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/out"
)

type CreateWarehouseInteractor struct {
	CreateWarehouseOutPort out.CreateWarehouseOutPort
}

func NewCreateWarehouseInteractor(
	CreateWarehouseOutPort out.CreateWarehouseOutPort,
) CreateWarehouseInteractor {
	return CreateWarehouseInteractor{CreateWarehouseOutPort}
}

func (s *CreateWarehouseInteractor) Execute(params *in.CreateWarehouseParams) (*entities.WarehouseEntity, error) {
	warehouse := entities.WarehouseEntity{
		Id:   uuid.New(),
		Name: params.Name,
	}
	createdErr := s.CreateWarehouseOutPort.CreateWarehouse(&warehouse)
	if createdErr != nil {
		return nil, createdErr
	}
	return &warehouse, nil
}
