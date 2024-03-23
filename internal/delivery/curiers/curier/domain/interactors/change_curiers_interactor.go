package interactors

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/out"
)

type ChangeCuriersStatusInteractor struct {
	FindCurierByIdWithOrdersPort out.FindCurierByIdWithOrdersOutPort
	UpdateCurierOutPort          out.UpdateCurierOutPort
}

func NewChangeCuriersStatusInteractor(
	FindCurierByIdWithOrdersPort out.FindCurierByIdWithOrdersOutPort,
	UpdateCurierOutPort out.UpdateCurierOutPort,
) ChangeCuriersStatusInteractor {
	return ChangeCuriersStatusInteractor{FindCurierByIdWithOrdersPort, UpdateCurierOutPort}
}

func (s *ChangeCuriersStatusInteractor) Execute(params *in.ChangeCuriersStatusParams) (*entities.CurierEntity, error) {
	curier := s.FindCurierByIdWithOrdersPort.FindCurierByIdWithOrders(params.CurierId)

	curier.ChangeStatus(params.IsActive)

	createdErr := s.UpdateCurierOutPort.UpdateCurier(&curier)
	if createdErr != nil {
		return nil, createdErr
	}
	return &curier, nil
}
