package orm

import (
	"github.com/google/uuid"
)

type OfferOrm struct {
	ID                     uuid.UUID `gorm:"primaryKey"`
	Name                   string
	OrderId                uuid.UUID
	ReportNumber           int
	VehicleType            string
	PreferredDeliveryAreas string
	WorkingHours           string
	Weight                 int
	Bid                    int
	CurierId               *uuid.UUID
}
