package dal

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

func CurierToOrm(curier *entities.CurierEntity) *orm.CurierOrm {
	var orders []orm.OrderOrm
	for _, order := range curier.Orders {
		orders = append(orders, orm.OrderOrm{
			ID:          order.Id,
			Name:        order.Name,
			Description: order.Description,
			OrderId:     order.OrderId,
			IsActive:    order.IsActive,
			CurierId:    order.CurierId,
			TotalSum:    order.TotalSum,
			Weight:      order.Weight,
		})
	}
	return &orm.CurierOrm{
		ID:               curier.Id,
		FirstName:        curier.FirstName,
		LastName:         curier.LastName,
		IsActive:         curier.IsActive,
		Email:            curier.Email,
		Phone:            curier.Phone,
		VehicleType:      curier.VehicleType,
		WorkingHours:     curier.WorkingHours,
		Rating:           curier.Rating,
		DeliveryCapacity: curier.DeliveryCapacity,
		Specialization:   curier.Specialization,
		CommissionRate:   curier.CommissionRate,
		PaymentDetails:   curier.PaymentDetails,
		Orders:           orders,
	}
}

func CurierToDomain(ormCurier *orm.CurierOrm) *entities.CurierEntity {
	var orders []entities.OrderEntity
	for _, order := range ormCurier.Orders {
		orders = append(orders, entities.OrderEntity{
			Id:          order.ID,
			Name:        order.Name,
			Description: order.Description,
			OrderId:     order.OrderId,
			IsActive:    order.IsActive,
			CurierId:    order.CurierId,
			TotalSum:    order.TotalSum,
			Weight:      order.Weight,
		})
	}
	return &entities.CurierEntity{
		Id:               ormCurier.ID,
		FirstName:        ormCurier.FirstName,
		LastName:         ormCurier.LastName,
		IsActive:         ormCurier.IsActive,
		Email:            ormCurier.Email,
		Phone:            ormCurier.Phone,
		VehicleType:      ormCurier.VehicleType,
		WorkingHours:     ormCurier.WorkingHours,
		Rating:           ormCurier.Rating,
		DeliveryCapacity: ormCurier.DeliveryCapacity,
		Specialization:   ormCurier.Specialization,
		CommissionRate:   ormCurier.CommissionRate,
		PaymentDetails:   ormCurier.PaymentDetails,
		Orders:           orders,
	}
}
