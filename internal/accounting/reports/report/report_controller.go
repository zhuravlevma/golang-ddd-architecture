package report

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/queries"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dtos"
)

type ReportController struct {
	createReportInteractor interactors.CreateReportInteractor
	updateReportInteractor interactors.UpdateReportInteractor
	findReportByIdQuery    queries.FindReportByIdQuery
}

func NewReportController(e *echo.Echo, createReportInteractor interactors.CreateReportInteractor, updateReportInteractor interactors.UpdateReportInteractor) *ReportController {
	controller := &ReportController{
		createReportInteractor: createReportInteractor,
		updateReportInteractor: updateReportInteractor,
	}
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
