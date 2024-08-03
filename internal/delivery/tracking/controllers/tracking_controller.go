package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/tracking/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/tracking/services"
)

type TrackingController struct {
	trackingService services.TrackingService
}

func NewVerificationController(e *echo.Echo, trackingService services.TrackingService) *TrackingController {
	controller := &TrackingController{
		trackingService: trackingService,
	}
	e.PATCH("/tracking", controller.UpdateRoute)

	return controller
}

func (rc *TrackingController) UpdateRoute(c echo.Context) error {
	var updateRouteDto dtos.UpdateRouteDto

	if err := c.Bind(&updateRouteDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := rc.trackingService.Update(&updateRouteDto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}
