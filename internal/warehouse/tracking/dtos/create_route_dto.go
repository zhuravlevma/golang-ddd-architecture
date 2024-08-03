package dtos

import "github.com/google/uuid"

type CreateRouteDto struct {
	CourierId uuid.UUID `json:"CourierId"`
	OrderId   uuid.UUID `json:"OrderId"`
}
