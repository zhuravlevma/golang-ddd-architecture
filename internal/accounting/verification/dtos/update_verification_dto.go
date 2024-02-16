package dtos

import "github.com/google/uuid"

type UpdateVerificationDto struct {
	ReportId  uuid.UUID `json:"reportId"`
	IsFull     *bool  `json:"isFull"`
	Description *string `json:"description"`
	ReportNumber *int `json:"ReportNumber"`
}
