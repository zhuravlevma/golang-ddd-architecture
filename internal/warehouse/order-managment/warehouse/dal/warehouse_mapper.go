package dal

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
)

func WarehouseToOrm(warehouse *entities.WarehouseEntity) *orm.WarehouseOrm {
	var orders []orm.OrderOrm
	for _, order := range warehouse.Orders {
		orders = append(orders, orm.OrderOrm{
			ID:          order.Id,
			Name:        order.Name,
			WarehouseId: warehouse.Id,
			IsValid:     order.IsValid,
		})
	}
	return &orm.WarehouseOrm{
		ID:     warehouse.Id,
		Name:   warehouse.Name,
		Orders: orders,
	}
}

func WarehouseToDomain(ormWarehouse *orm.WarehouseOrm) *entities.WarehouseEntity {
	var orders []entities.OrderEntity
	for _, order := range ormWarehouse.Orders {
		orders = append(orders, entities.OrderEntity{
			Id:      order.ID,
			Name:    order.Name,
			IsValid: order.IsValid,
		})
	}
	return &entities.WarehouseEntity{
		Id:     ormWarehouse.ID,
		Name:   ormWarehouse.Name,
		Orders: orders,
	}
}
