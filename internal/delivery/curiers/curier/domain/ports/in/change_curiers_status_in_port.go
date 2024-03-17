package in

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type ChangeCuriersStatusParams struct {
	CurierId uuid.UUID
	IsActive bool
}

type ChangeCuriersStatusInPort interface {
	Execute(changeCuriersStatusParams *ChangeCuriersStatusParams) (*entities.CurierEntity, error)
}
