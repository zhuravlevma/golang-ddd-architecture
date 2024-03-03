package events

import (
	"github.com/google/uuid"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

type OrderValidatedPayload struct {
	OrderId uuid.UUID `json:"OrderId"`
}

type OrderValidatedEvent struct {
	Event lib.DomainAttr
}

func (e *OrderValidatedEvent) GetEvent() *lib.DomainAttr {
	return &e.Event
}

func NewOrderValidatedEvent(payload OrderValidatedPayload, aggregateId uuid.UUID) *OrderValidatedEvent {
	return &OrderValidatedEvent{
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
