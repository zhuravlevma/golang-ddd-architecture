package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	amqp "github.com/rabbitmq/amqp091-go"
	config "github.com/zhuravlevma/golang-ddd-architecture/internal/__config__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/interactors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := config.New()
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Db.Host, config.Db.User, config.Db.Password, config.Db.Name, config.Db.Port, config.Db.SSL)
	port := fmt.Sprintf(":%d", config.Port)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer func() {
		log.Print("close")
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close()
	}()

	gormDB, err := gorm.Open(postgres.Open(url), &gorm.Config{})
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

	report.NewReportController(e, ch, config, createReportService, updateReportService)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
