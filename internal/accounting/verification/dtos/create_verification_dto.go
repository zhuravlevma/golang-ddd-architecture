package dtos

import "github.com/google/uuid"

type CreateVerificationDto struct {
	ReportId     uuid.UUID `json:"reportId"`
	IsFull       bool      `json:"isFull"`
	Description  string    `json:"description"`
	ReportNumber int       `json:"ReportNumber"`
}
