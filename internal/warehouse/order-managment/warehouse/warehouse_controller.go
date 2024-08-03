package curier

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rabbitmq/amqp091-go"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/dtos"
)

type WarehouseController struct {
	AddOrderInteractor        in.AddOrderInPort
	CreateWarehouseInteractor in.CreateWarehouseInPort
	UpdateOrderInteractor     in.UpdateOrderInPort
}

func NewWarehouseController(e *echo.Echo, amqpChannel *amqp091.Channel, config *config.Config,
	AddOrderInteractor interactors.AddOrderInteractor,
	CreateWarehouseInteractor interactors.CreateWarehouseInteractor,
	UpdateOrderInteractor interactors.UpdateOrderInteractor,
) *WarehouseController {
	controller := &WarehouseController{
		AddOrderInteractor:        &AddOrderInteractor,
		CreateWarehouseInteractor: &CreateWarehouseInteractor,
		UpdateOrderInteractor:     &UpdateOrderInteractor,
	}

	e.POST("/warehouse/", controller.CreateWarehouse)

	return controller
}

func (oc *WarehouseController) CreateWarehouse(c echo.Context) error {
	var createWarehouseDto dtos.CreateWarehouseDto

	if err := c.Bind(&createWarehouseDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.CreateWarehouseInteractor.Execute(&in.CreateWarehouseParams{
		Name: createWarehouseDto.Name,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *WarehouseController) AddOrderToWh(c echo.Context) error {
	var addOrderDto dtos.AddOrderDto

	id, _ := uuid.Parse(c.Param("id"))

	if err := c.Bind(&addOrderDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.AddOrderInteractor.Execute(&in.AddOrderParams{
		WarehouseId: id,
		OrderId:     addOrderDto.OrderId,
		Name:        addOrderDto.Name,
		IsValid:     addOrderDto.IsValid,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func (oc *WarehouseController) UpdateOrder(c echo.Context) error {
	var updateOrderDto dtos.UpdateOrderDto

	id, _ := uuid.Parse(c.Param("id"))

	orderId, _ := uuid.Parse(c.Param("orderId"))

	if err := c.Bind(&updateOrderDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := oc.UpdateOrderInteractor.Execute(&in.UpdateOrderParams{
		IsValid:     updateOrderDto.IsValid,
		WarehouseId: id,
		OrderId:     orderId,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}
