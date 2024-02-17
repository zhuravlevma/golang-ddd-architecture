package entities

import "github.com/google/uuid"

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
	CommissionRate   int
	PaymentDetails   int
	Orders           []OrderEntity
}
