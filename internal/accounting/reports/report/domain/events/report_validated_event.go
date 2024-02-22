package events

import (
	"github.com/google/uuid"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
)

type ReportValidatedPayload struct {
	OrderId uuid.UUID `json:"OrderId"`
}

type ReportValidatedEvent struct {
	Reason        string                `json:"Reason"`
	Payload       ReportValidatedPayload `json:"Payload"`
	MessageName   string                `json:"MessageName"`
	AggregateId   uuid.UUID                `json:"AggregateId"`
	AggregateName string                `json:"AggregateName"`
	ContextName   string                `json:"ContextName"`
	MessageType   string                `json:"MessageType"`
}

func NewReportValidatedEvent(payload ReportValidatedPayload, aggregateId uuid.UUID) *ReportValidatedEvent {
	return &ReportValidatedEvent{
		Reason:        "The report was validated",
		Payload:       payload,
		MessageName:   config.New().ReportValidatedEvent,
		AggregateId:   aggregateId,
		AggregateName: "Report",
		ContextName:   "accounting",
	}
}
