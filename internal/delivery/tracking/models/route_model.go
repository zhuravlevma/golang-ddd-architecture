package models

import (
	"github.com/google/uuid"
)

type Route struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	StartLatitude  *string
	StartLongitude *string
	EndLatitude    *string
	EndLongitude   *string
	Completed      bool
	OrderId        uuid.UUID
	CourierId      uuid.UUID
}
