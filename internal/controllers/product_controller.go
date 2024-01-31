package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/services"

	"github.com/google/uuid"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(e *echo.Echo, service services.ProductService) *ProductController {
	controller := &ProductController{
		service: service,
	}
	e.GET("/products/:id", controller.GetProductByID)

	return controller
}

func (pc *ProductController) GetProductByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID format",
		})
	}

	product, err := pc.service.FindProductByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch product",
		})
	}

	return c.JSON(http.StatusOK, product)
}
