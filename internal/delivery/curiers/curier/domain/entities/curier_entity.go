package entities

import (
	"errors"

	"github.com/google/uuid"
)

type CurierEntity struct {
	Id               uuid.UUID
	FirstName        string
	LastName         string
	IsActive         bool
	Email            string
	Phone            int
	VehicleType      string
	WorkingHours     int
	Rating           float64
	DeliveryCapacity int
	Specialization   string
	CommissionRate   float64
	PaymentDetails   int
	Orders           []OrderEntity
}

func (c *CurierEntity) AddOrder(order OrderEntity) error {
	if len(c.Orders) > 3 {
		return errors.New("Exceeded the number of orders");
	}
	c.Orders = append(c.Orders, order)
	return nil
}

func (c *CurierEntity) ChangeFirstName(firstName string) {
	c.FirstName = firstName
}

func (c *CurierEntity) ChangeLastName(lastName string) {
	c.LastName = lastName
}

func (c *CurierEntity) UpdateRating(newRating float64) {
	totalRating := c.Rating * float64(len(c.Orders))
	updatedRating := (totalRating + newRating) / (float64(len(c.Orders)) + 1);
	c.Rating = updatedRating
}

func (c *CurierEntity) SetDeliveryCapacity(capacity int) error {
	for _, order := range c.Orders {
		if order.Weight > capacity {
			return errors.New("Delivery capacity is insufficient for existing orders.")
		}
	}

	c.DeliveryCapacity = capacity
	return nil
}

func (c *CurierEntity) changeSpecialization(area string) {
	c.Specialization = area
}

func (c *CurierEntity) SetCommissionRate(rate float64) error {
	if rate > 0.5 {
		return errors.New("Commission rate cannot exceed 50%.")
	}
	c.CommissionRate = rate
	return nil
}

func (c *CurierEntity) UpdatePaymentDetails(bankAddress int) error {
	for _, order := range c.Orders {
		if order.IsActive == false {
			return errors.New("Cannot update payment details for orders with pending payment.")
		}
	}
	c.PaymentDetails = bankAddress
	return nil
}

func (c *CurierEntity) DeliverOrder(orderId uuid.UUID) {
	for _, order := range c.Orders {
		if order.Id == orderId {
			order.DeliverOrder()
		}
	}
}


func (c *CurierEntity) CompleteDeliveryForAllOrders() {
	for _, order := range c.Orders {
		order.DeliverOrder()
	}
}

func (c *CurierEntity) ChangeStatus(newStatus bool) error {
	if c.IsActive == true && newStatus == false && len(c.Orders) > 0 {
		errors.New("Deliverman has orders");
	}
	c.IsActive = newStatus
	return nil
}

func (c *CurierEntity) AddNewMessageToOrders(message string) {
	for _, order := range c.Orders {
		order.AddInfoToDescription(message + " " + c.FirstName+ " " + c.LastName)
	}
}
