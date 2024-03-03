package entities

import (
	"github.com/google/uuid"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/events"
)

type WarehouseEntity struct {
	Id uuid.UUID
	Name string
	Orders []OrderEntity
	DomainMessages []lib.DomainMessage
}

func (w *WarehouseEntity) AddOrder(order OrderEntity) {
	w.Orders = append(w.Orders, order)
}

func (w *WarehouseEntity) ChangeOrderStatusToValid(orderId uuid.UUID) {
	for _, order := range w.Orders {
		if order.Id == orderId {
			order.ChangeStatus(true)
			w.DomainMessages = append(w.DomainMessages, events.NewOrderValidatedEvent(events.OrderValidatedPayload{
				OrderId: orderId,
			}, w.Id))
		}
	}

}
