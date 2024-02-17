package dtos

import "github.com/google/uuid"

type UpdateVerificationDto struct {
	Id        uuid.UUID `json:"Id"`
	IsFull    *bool     `json:"isFull"`
	Completed *bool     `json:"isCompleted"`
	Signed    *bool     `json:"Signed"`
}
