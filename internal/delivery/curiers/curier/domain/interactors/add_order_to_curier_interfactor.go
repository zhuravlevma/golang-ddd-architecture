package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/out"
)

type AddOrderToCurierInteractor struct {
	FindCurierByIdWithOrdersPort out.FindCurierByIdWithOrdersOutPort
	UpdateCurierOutPort          out.UpdateCurierOutPort
}

func NewCreateReportInteractor(
	FindCurierByIdWithOrdersPort out.FindCurierByIdWithOrdersOutPort,
	UpdateCurierOutPort out.UpdateCurierOutPort,
) AddOrderToCurierInteractor {
	return AddOrderToCurierInteractor{FindCurierByIdWithOrdersPort, UpdateCurierOutPort}
}

func (s *AddOrderToCurierInteractor) Execute(params *in.AddOrderToCurierParams) (*entities.CurierEntity, error) {
	curier := s.FindCurierByIdWithOrdersPort.FindCurierByIdWithOrders(params.CurierId)

	curier.AddOrder(entities.OrderEntity{
		Id:          uuid.New(),
		Name:        "test",
		Description: "test desc",
		IsActive:    false,
		OrderId:     params.OrderId,
		TotalSum:    0,
		Weight:      1,
		CurierId:    params.CurierId,
	})
	createdErr := s.UpdateCurierOutPort.UpdateCurier(&curier)
	if createdErr != nil {
		return nil, createdErr
	}
	return &curier, nil
}
