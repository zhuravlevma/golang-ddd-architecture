package events

import (
	"github.com/google/uuid"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

type OrderValidatedPayload struct {
	OrderId uuid.UUID `json:"OrderId"`
}

func NewOrderValidatedEvent(payload OrderValidatedPayload, aggregateId uuid.UUID) lib.DomainMessage[lib.DomainMessagePayload] {
	return lib.DomainMessage[lib.DomainMessagePayload]{
		Reason:        "The report was validated",
		Payload:       payload,
		MessageName:   config.New().ReportValidatedEvent,
		AggregateId:   aggregateId,
		AggregateName: "Report",
		ContextName:   "accounting",
	}
}
