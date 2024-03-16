package relay

import (
	"github.com/jackc/pgx/pgtype"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

func DomainToORM(event lib.DomainMessage[lib.DomainMessagePayload]) *MessageOrm {
	return &MessageOrm{
		Reason:        event.Reason,
		MessageType:   event.MessageType,
		Payload:       event.Payload.(pgtype.JSONB),
		AggregateId:   event.AggregateId,
		AggregateName: event.AggregateName,
		ContextName:   event.ContextName,
		MessageName:   event.MessageName,
	}
}
