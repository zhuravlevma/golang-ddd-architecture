package out

import "github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"

type UpdateCurierOutPort interface {
	UpdateCurier(*entities.CurierEntity) error
}
