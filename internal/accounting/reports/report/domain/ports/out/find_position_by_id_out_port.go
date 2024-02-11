package out

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/entities"
)
type FindPositionByIdOutPort interface {
	findPositionById(orderId uuid.UUID) (*entities.ReportPositionEntity, error)
}
