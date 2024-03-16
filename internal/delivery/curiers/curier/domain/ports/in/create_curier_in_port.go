package in

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
)

type CreateCurierParams struct {
	FirstName string
  LastName string
}

type CreateCurierInPort interface {
	Execute(createCurierParams *CreateCurierParams) (*entities.CurierEntity, error)
}
