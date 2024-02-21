package config

import (
	"os"
	"strconv"
	"strings"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
	SSL      string
}

type AmqpConfig struct {
	Host     string
	User     string
	Password string
	Port     int
}

type Config struct {
	Db                   DbConfig
	Amqp                 AmqpConfig
	Port                 int
	ReportValidatedEvent string
	OfferTakedEvent      string
	OrderValidatedEvent  string
}

func New() *Config {
	return &Config{
		Db: DbConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("postgres", "postgres"),
			Name:     getEnv("DB_NAME", "postgres"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			SSL:      getEnv("DB_SSL_MODE", "disable"),
		},
		Amqp: AmqpConfig{
			Host:     getEnv("AMQP_HOST", "localhost"),
			User:     getEnv("AMQP_USER", "guest"),
			Password: getEnv("AMQP_PASSWORD", "guest"),
			Port:     getEnvAsInt("AMQP_PORT", 5672),
		},
		Port:                 getEnvAsInt("PORT", 8080),
		ReportValidatedEvent: getEnv("REPORT_VALIDATED_EVENT", "REPORT_VALIDATED"),
		OfferTakedEvent:      getEnv("OFFER_TAKED_EVENT", "OFFER_TAKED"),
		OrderValidatedEvent:  getEnv("ORDER_VALIDATED_EVENT", "ORDER_VALIDATED"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
