package offer

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/interactors"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/dtos"
)

type OfferController struct {
	UpdateOfferInteractor in.UpdateOfferInPort
	CreateOfferInteractor in.CreateOfferInPort
}

func NewOfferController(e *echo.Echo, updateOfferInteractor interactors.UpdateOfferInteractor, createOfferInteractor interactors.CreateOfferInteractor) *OfferController {
	controller := &OfferController{
		UpdateOfferInteractor: &updateOfferInteractor,
		CreateOfferInteractor: &createOfferInteractor,
	}
	e.PATCH("/offers/:id", controller.UpdateOrderStatus)

	return controller
}

func (oc *OfferController) UpdateOrderStatus(c echo.Context) error {
	var updateReportDto dtos.UpdateOfferDto

	id, err := uuid.Parse(c.Param("id"))
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
