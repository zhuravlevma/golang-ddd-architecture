package events

import (
	"github.com/google/uuid"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

type OfferTakedPayload struct {
	OrderId uuid.UUID `json:"OrderId"`
}

func NewOfferTakedEvent(payload OfferTakedPayload, aggregateId uuid.UUID) lib.DomainMessage[lib.DomainMessagePayload] {
	return lib.DomainMessage[lib.DomainMessagePayload]{
		Reason:        "The delivery man accepted the offer",
		Payload:       payload,
		MessageName:   config.New().OfferTakedEvent,
		AggregateId:   aggregateId,
		AggregateName: "Offer",
		ContextName:   "deivery",
	}
}
