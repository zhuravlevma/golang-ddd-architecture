package lib

import "github.com/google/uuid"

type DomainMessagePayload interface{}
type DomainMessage[T DomainMessagePayload] struct {
	Reason        string    `json:"Reason"`
	Payload       T         `json:"Payload"`
	MessageName   string    `json:"MessageName"`
	AggregateId   uuid.UUID `json:"AggregateId"`
	AggregateName string    `json:"AggregateName"`
	ContextName   string    `json:"ContextName"`
	MessageType   string    `json:"MessageType"`
}
