package orm

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/pgtype"
)

type MessageOrm struct {
	ID            int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	MessageType   string
	Reason        string
	Payload       pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Published     bool
	MessageName   string
	AggregateId   uuid.UUID
	AggregateName string
	ContextName   string
}
