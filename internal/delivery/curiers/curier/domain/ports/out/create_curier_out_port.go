package out

import "github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"

type CreateCurierOutPort interface {
	CreateCurier(*entities.CurierEntity) error
}
