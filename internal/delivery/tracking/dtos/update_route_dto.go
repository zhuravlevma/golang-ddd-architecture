package dtos

import "github.com/google/uuid"

type UpdateRouteDto struct {
	Id        uuid.UUID  `json:"Id"`
	CourierId *uuid.UUID `json:"CourierId"`
	OrderId   *uuid.UUID `json:"OrderId"`
}
