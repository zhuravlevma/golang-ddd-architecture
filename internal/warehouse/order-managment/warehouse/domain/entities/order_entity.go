package entities

import "github.com/google/uuid"

type OrderEntity struct {
	Id      uuid.UUID
	Name    string
	IsValid bool
}

func (o *OrderEntity) ChangeStatus(isValid bool) {
	o.IsValid = isValid
}
