package orm

import (
	"github.com/jackc/pgx/pgtype"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
)

func DomainToORM(event lib.DomainMessage) *MessageOrm {
	eventAttr := event.GetEvent()
	return &MessageOrm{
		Reason:        eventAttr.Reason,
		MessageType:   eventAttr.MessageType,
		Payload:       eventAttr.Payload.(pgtype.JSONB),
		AggregateId:   eventAttr.AggregateId,
		AggregateName: eventAttr.AggregateName,
		ContextName:   eventAttr.ContextName,
		MessageName:   eventAttr.MessageName,
	}
}
