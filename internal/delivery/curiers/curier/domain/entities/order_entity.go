package entities

import (
	"errors"

	"github.com/google/uuid"
)

type OrderEntity struct {
	Id          uuid.UUID
	Name        string
	Description string
	IsActive    bool
	Weight      int
	TotalSum    int
	CurierId    uuid.UUID
	OrderId     uuid.UUID
}

func (o *OrderEntity) checkName() error {
	if len(o.Name) < 3 {
		return errors.New("The length of the name is less than 3")
	}
	return nil
}

func (o *OrderEntity) MarkAsDelayedDueToTraffic() {
	o.IsActive = false
	o.AddInfoToDescription("Order delayed due to heavy traffic.")
}

func (o *OrderEntity) RequestGiftWrapping() {
	o.AddInfoToDescription("Gift wrapping requested.")
	o.TotalSum += 5
}

func (o *OrderEntity) CancelOrder() error {
	if o.IsActive == true {
		o.IsActive = false
		o.AddInfoToDescription("Order cancelled by customer.")
	} else {
		return errors.New("Order cannot be cancelled. Invalid order status.")
	}
	return nil
}

func (o *OrderEntity) ApplyTip(tipAmount int) error {
	if o.IsActive == true {
		o.TotalSum += tipAmount
		o.AddInfoToDescription("Tip applied")
	} else {
		return errors.New("Tip cannot be applied. Order is not delivered.")
	}
	return nil
}

func (o *OrderEntity) DeliverOrder() {
	o.IsActive = false
	o.AddInfoToDescription("This order has been delivered.")
}

func (o *OrderEntity) ReturnOrder() {
	o.IsActive = false
	o.AddInfoToDescription("This order has been returned :(")
}

func (o *OrderEntity) AddInfoToDescription(info string) {
	o.Description += "\n" + info
}
