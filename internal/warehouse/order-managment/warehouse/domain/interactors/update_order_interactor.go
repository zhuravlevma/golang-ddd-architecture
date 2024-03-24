package interactors

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/out"
)

type UpdateOrderInteractor struct {
	GetWarehouseWithOrderPort out.GetWarehouseWithOrderOutPort
	UpdateWarehousePort          out.UpdateWarehouseOutPort
}

func NewUpdateOrderInteractor(
	GetWarehouseWithOrderPort out.GetWarehouseWithOrderOutPort,
	UpdateWarehousePort out.UpdateWarehouseOutPort,
) UpdateOrderInteractor {
	return UpdateOrderInteractor{GetWarehouseWithOrderPort, UpdateWarehousePort}
}

func (s *UpdateOrderInteractor) Execute(params *in.UpdateOrderParams) (*entities.WarehouseEntity, error) {
	warehouse := s.GetWarehouseWithOrderPort.GetWarehouseWithOrder(params.WarehouseId, params.OrderId)

	if params.IsValid {
		warehouse.ChangeOrderStatusToValid(params.OrderId)
	}

	createdErr := s.UpdateWarehousePort.UpdateWarehouse(&warehouse)
	if createdErr != nil {
		return nil, createdErr
	}
	return &warehouse, nil
}
