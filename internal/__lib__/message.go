package lib

import "github.com/google/uuid"

type DomainMessagePayload interface{}
type DomainAttr struct {
	Reason        string               `json:"Reason"`
	Payload       DomainMessagePayload `json:"Payload"`
	MessageName   string               `json:"MessageName"`
	AggregateId   uuid.UUID            `json:"AggregateId"`
	AggregateName string               `json:"AggregateName"`
	ContextName   string               `json:"ContextName"`
	MessageType   string               `json:"MessageType"`
}

type DomainMessage interface {
	GetEvent() *DomainAttr
}
