package interactors

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/entities"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/out"
)

type CreateCurierInteractor struct {
	CreateCurierOutPort out.CreateCurierOutPort
}

func NewCreateCurierInteractor(
	CreateCurierOutPort out.CreateCurierOutPort,
) CreateCurierInteractor {
	return CreateCurierInteractor{CreateCurierOutPort}
}

func (s *CreateCurierInteractor) Execute(params *in.CreateCurierParams) (*entities.CurierEntity, error) {
	curier := entities.CurierEntity{
		Id:               uuid.New(),
		FirstName:        params.FirstName,
		LastName:         params.LastName,
		Email:            "email",
		Phone:            121212,
		VehicleType:      "bike",
		WorkingHours:     10,
		Rating:           0,
		DeliveryCapacity: 2,
		Specialization:   "food",
		CommissionRate:   0.2,
		PaymentDetails:   93321331332,
		IsActive:         true,
	}

	createdErr := s.CreateCurierOutPort.CreateCurier(&curier)

	if createdErr != nil {
		return nil, createdErr
	}
	return &curier, nil
}
