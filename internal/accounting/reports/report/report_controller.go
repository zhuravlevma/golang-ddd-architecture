package report

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	amqp "github.com/rabbitmq/amqp091-go"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/queries"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/events"
)

type ReportController struct {
	createReportInteractor interactors.CreateReportInteractor
	updateReportInteractor interactors.UpdateReportInteractor
	findReportByIdQuery    queries.FindReportByIdQuery
}

func NewReportController(e *echo.Echo, amqpChannel *amqp.Channel, config *config.Config, createReportInteractor interactors.CreateReportInteractor, updateReportInteractor interactors.UpdateReportInteractor) *ReportController {
	controller := &ReportController{
		createReportInteractor: createReportInteractor,
		updateReportInteractor: updateReportInteractor,
	}
	q, err := amqpChannel.QueueDeclare(
		config.OrderValidatedEvent,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	messages, err := amqpChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	go func() {
		for message := range messages {
			data := &events.OrderValidatedEvent{}
			err := json.Unmarshal(message.Body, data)
			if err != nil {
				panic(err)
			}
			controller.ApplyOrderValidated(data)
		}
	}()
	e.PATCH("/reports/:id", controller.UpdateReport)
	e.GET("/reports/:id", controller.FindReportById)

	return controller
}

func (rc *ReportController) UpdateReport(c echo.Context) error {
	var updateReportDto dtos.UpdateReportDto

	id, err := uuid.Parse(c.Param("id"))
	if err := c.Bind(&updateReportDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := rc.updateReportInteractor.Execute(&in.UpdateReportParams{
		ReportId: id,
		IsValid:  &updateReportDto.IsValid,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (rc *ReportController) FindReportById(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))

	result, err := rc.findReportByIdQuery.Execute(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (rc *ReportController) ApplyOrderValidated(event *events.OrderValidatedEvent) error {
	_, err := rc.createReportInteractor.Execute(&in.CreateReportParams{
		OrderId: event.Payload.OrderId,
	})

	if err != nil {
		return err
	}
	return nil
}
