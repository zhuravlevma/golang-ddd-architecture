package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/interactors"

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

	gormDB.AutoMigrate(&orm.ReportOrm{}, &orm.ReportPositionOrm{})

	reportRepository := dal.ReportRepository{
		Db: gormDB,
	}

	createReportService := interactors.NewCreateReportInteractor(&reportRepository)
	updateReportService := interactors.NewUpdateReportInteractor(&reportRepository, &reportRepository)

	e := echo.New()

	report.NewReportController(e, createReportService, updateReportService)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
