package offer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rabbitmq/amqp091-go"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	lib "github.com/zhuravlevma/golang-ddd-architecture/internal/__lib__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/events"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/dtos"
)

type OfferController struct {
	UpdateOfferInteractor in.UpdateOfferInPort
	CreateOfferInteractor in.CreateOfferInPort
}

func NewOfferController(e *echo.Echo, amqpChannel *amqp091.Channel, config *config.Config, updateOfferInteractor interactors.UpdateOfferInteractor, createOfferInteractor interactors.CreateOfferInteractor) *OfferController {
	controller := &OfferController{
		UpdateOfferInteractor: &updateOfferInteractor,
		CreateOfferInteractor: &createOfferInteractor,
	}

	messages, err := amqpChannel.Consume(
		config.ReportValidatedEvent, // queue
		"",                          // consumer
		true,                        // auto-ack
		false,                       // exclusive
		false,                       // no-local
		false,                       // no-wait
		nil,                         // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	go func() {
		for message := range messages {
			controller.ApplyOrderValidated(message)
		}
	}()
	e.PATCH("/offers/:id", controller.UpdateOrderStatus)

	return controller
}

func (oc *OfferController) UpdateOrderStatus(c echo.Context) error {
	var updateReportDto dtos.UpdateOfferDto

	id, _ := uuid.Parse(c.Param("id"))
	if err := c.Bind(&updateReportDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.UpdateOfferInteractor.Execute(&in.UpdateOfferParams{
		OfferId:  id,
		CurierId: updateReportDto.CurierId,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *OfferController) ApplyOrderValidated(message amqp091.Delivery) error {
	event := &lib.DomainMessage[events.ReportValidatedPayload]{}
	err := json.Unmarshal(message.Body, event)
	if err != nil {
		panic(err)
	}

	_, err = oc.CreateOfferInteractor.Execute(&in.CreateOfferParams{
		OrderId: event.Payload.OrderId,
		Name:    "Report with " + event.Payload.OrderId.String(),
	})

	if err != nil {
		return err
	}
	return nil
}
