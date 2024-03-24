package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/out"
)

type AddOrderInteractor struct {
	GetWarehouseWithOrderPort out.GetWarehouseWithOrderOutPort
	UpdateWarehousePort          out.UpdateWarehouseOutPort
}

func NewCreateReportInteractor(
	GetWarehouseWithOrderPort out.GetWarehouseWithOrderOutPort,
	UpdateWarehousePort out.UpdateWarehouseOutPort,
) AddOrderInteractor {
	return AddOrderInteractor{GetWarehouseWithOrderPort, UpdateWarehousePort}
}

func (s *AddOrderInteractor) Execute(params *in.AddOrderParams) (*entities.WarehouseEntity, error) {
	warehouse := s.GetWarehouseWithOrderPort.GetWarehouseWithOrder(params.WarehouseId, params.OrderId)

	warehouse.AddOrder(entities.OrderEntity{
		Id:          uuid.New(),
		Name:        params.Name,
		IsValid: false,
	})
	createdErr := s.UpdateWarehousePort.UpdateWarehouse(&warehouse)
	if createdErr != nil {
		return nil, createdErr
	}
	return &warehouse, nil
}
