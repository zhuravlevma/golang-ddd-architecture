package dtos

import "github.com/google/uuid"

type UpdateOfferDto struct {
	CurierId *uuid.UUID `json:"CurierId"`
}
