package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/controllers"
	psql "github.com/zhuravlevma/golang-ddd-architecture/internal/test/db/postgres"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/test/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {
	dsn := "host=localhost user=maksim password=postgres dbname=postgres port=5432 sslmode=disable"
	port := ":8080"

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gormDB.AutoMigrate(&psql.Product{})

	productRepo := psql.NewGormProductRepository(gormDB)

	productService := services.NewProductService(productRepo)

	e := echo.New()

	controllers.NewProductController(e, productService)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
