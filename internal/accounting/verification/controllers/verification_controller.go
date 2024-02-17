package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/services"
)

type VerificationController struct {
	verificationService services.VerificationService
}

func NewVerificationController(e *echo.Echo, verificationService services.VerificationService) *VerificationController {
	controller := &VerificationController{
		verificationService: verificationService,
	}
	e.PATCH("/verification", controller.UpdateVerification)

	return controller
}

func (rc *VerificationController) UpdateVerification(c echo.Context) error {
	var updateReportDto dtos.UpdateVerificationDto

	if err := c.Bind(&updateReportDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := rc.verificationService.Update(&updateReportDto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
}
