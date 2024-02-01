package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/controllers/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/services"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(e *echo.Echo, service services.ProductService) *ProductController {
	controller := &ProductController{
		service: service,
	}
	e.GET("/products/:id", controller.GetProductByID)
	e.POST("/products", controller.CreateProduct)

	return controller
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var createProductDto dtos.CreateProductDto

	if err := c.Bind(&createProductDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := pc.service.CreateProduct(&createProductDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, result)
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
