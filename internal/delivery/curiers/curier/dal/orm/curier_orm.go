package orm

import (
	"github.com/google/uuid"
)

type CurierOrm struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	FirstName      string
	LastName      string
	IsActive bool
	Email string
	Phone int
	VehicleType string
	WorkingHours int
	Rating int
	DeliveryCapacity int
	Specialization string
	CommissionRate int
	PaymentDetails int
	Orders    []OrderOrm `gorm:"foreignKey:CurierId"`
}
