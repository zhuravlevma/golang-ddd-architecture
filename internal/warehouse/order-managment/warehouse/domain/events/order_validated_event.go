package events

import (
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
)

type OrderValidatedPayload struct {
	orderId string
}

type OrderValidatedEvent struct {
	Reason        string                `json:"Reason"`
	Payload       OrderValidatedPayload `json:"Payload"`
	MessageName   string                `json:"MessageName"`
	AggregateId   string                `json:"AggregateId"`
	AggregateName string                `json:"AggregateName"`
	ContextName   string                `json:"ContextName"`
	MessageType   string                `json:"MessageType"`
}

func (e *OrderValidatedEvent) New(payload OrderValidatedPayload, aggregateId string) *OrderValidatedEvent {
	return &OrderValidatedEvent{
		Reason:        "The order was validated",
		Payload:       payload,
		MessageName:   config.New().OrderValidatedEvent,
		AggregateId:   aggregateId,
		AggregateName: "Warehouse",
		ContextName:   "warehouse",
	}
}
