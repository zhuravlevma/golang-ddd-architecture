package events

import (
	"github.com/google/uuid"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

type ReportValidatedPayload struct {
	OrderId uuid.UUID `json:"OrderId"`
}

type ReportValidatedEvent struct {
	Event lib.DomainAttr
}

func (e *ReportValidatedEvent) GetEvent() *lib.DomainAttr {
	return &e.Event
}

func NewReportValidatedEvent(payload ReportValidatedPayload, aggregateId uuid.UUID) *ReportValidatedEvent {
	return &ReportValidatedEvent{
		Event: lib.DomainAttr{
			Reason:        "The report was validated",
			Payload:       payload,
			MessageName:   config.New().ReportValidatedEvent,
			AggregateId:   aggregateId,
			AggregateName: "Report",
			ContextName:   "accounting",
		},
	}
}
