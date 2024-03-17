package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type UpdateCuriersInfoParams struct {
	CurierId  uuid.UUID
	FirstName *string
	LastName  *string
	IsActive  *bool
}

type UpdateCuriersInPort interface {
	Execute(updateCuriersInfoParams *UpdateCuriersInfoParams) (*entities.CurierEntity, error)
}
