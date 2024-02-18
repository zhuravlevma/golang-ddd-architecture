package dtos

import "github.com/google/uuid"

type AddOrderToCurierDto struct {
	CurierId uuid.UUID `json:"CurierId"`
	OrderId uuid.UUID `json:"OrderId"`
}
