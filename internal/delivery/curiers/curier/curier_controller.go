package curier

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rabbitmq/amqp091-go"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/events"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/curiers/curier/dtos"
)

type CurierController struct {
	CreateCurierInteractor        in.CreateCurierInPort
	AddOrderToCurierInteractor    in.AddOrderToCurierInPort
	ChangeCuriersStatusInteractor in.ChangeCuriersStatusInPort
}

func NewCurierController(e *echo.Echo, amqpChannel *amqp091.Channel, config *config.Config,
	CreateCurierInteractor interactors.CreateCurierInteractor,
	AddOrderToCurierInteractor interactors.AddOrderToCurierInteractor,
	ChangeCuriersStatusInteractor interactors.ChangeCuriersStatusInteractor,
) *CurierController {
	controller := &CurierController{
		CreateCurierInteractor:        &CreateCurierInteractor,
		AddOrderToCurierInteractor:    &AddOrderToCurierInteractor,
		ChangeCuriersStatusInteractor: &ChangeCuriersStatusInteractor,
	}

	messages, err := amqpChannel.Consume(
		config.OfferTakedEvent, // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	go func() {
		for message := range messages {
			controller.ApplyOfferTaked(message)
		}
	}()
	e.POST("/curiers/", controller.CreateCurier)

	return controller
}

func (oc *CurierController) CreateCurier(c echo.Context) error {
	var createCurierDto dtos.CreateCurierDto

	if err := c.Bind(&createCurierDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.CreateCurierInteractor.Execute(&in.CreateCurierParams{
		FirstName: createCurierDto.FirstName,
		LastName:  createCurierDto.LastName,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *CurierController) AddOrderToCurier(c echo.Context) error {
	var addOrderToCurierDto dtos.AddOrderToCurierDto

	if err := c.Bind(&addOrderToCurierDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.AddOrderToCurierInteractor.Execute(&in.AddOrderToCurierParams{
		OrderId:  addOrderToCurierDto.OrderId,
		CurierId: addOrderToCurierDto.CurierId,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *CurierController) ChangeCuriersStatus(c echo.Context) error {
	var changeCuriersStatusDto dtos.ChangeCuriersStatusDto

	id, _ := uuid.Parse(c.Param("id"))

	if err := c.Bind(&changeCuriersStatusDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.ChangeCuriersStatusInteractor.Execute(&in.ChangeCuriersStatusParams{
		IsActive: changeCuriersStatusDto.IsActive,
		CurierId: id,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *CurierController) ApplyOfferTaked(message amqp091.Delivery) error {
	event := &lib.DomainMessage[events.OfferTakedPayload]{}
	err := json.Unmarshal(message.Body, event)
	if err != nil {
		panic(err)
	}

	_, err = oc.AddOrderToCurierInteractor.Execute(&in.AddOrderToCurierParams{
		OrderId:  event.Payload.OrderId,
		CurierId: event.Payload.CurierId,
	})

	if err != nil {
		return err
	}
	return nil
}
