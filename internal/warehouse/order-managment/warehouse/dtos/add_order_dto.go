package dtos

import "github.com/google/uuid"

type AddOrderDto struct {
	Name    string    `json:"Name"`
	OrderId uuid.UUID `json:"OrderId"`
	IsValid bool      `json:"IsValid"`
}
