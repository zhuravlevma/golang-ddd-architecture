package config

import (
	"os"
	"strconv"
	"strings"
)

type DbConfig struct {
    Host string
    User   string
		Password   string
		Name string
		Port int
		SSL string
}

type Config struct {
    Db    DbConfig
    Port int
}

// New returns a new Config struct
func New() *Config {
    return &Config{
			Db: DbConfig{
				Host: getEnv("DB_HOST", "localhost"),
				User: getEnv("DB_USER", "postgres"),
				Password: getEnv("postgres", "postgres"),
				Name: getEnv("DB_NAME", "postgres"),
				Port: getEnvAsInt("DB_PORT", 5432),
				SSL: getEnv("DB_SSL_MODE", "disable"),
			},
			Port: getEnvAsInt("PORT", 8080),
    }
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
	return value
    }

    return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
    valStr := getEnv(name, "")
    if val, err := strconv.ParseBool(valStr); err == nil {
	return val
    }

    return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
    valStr := getEnv(name, "")

    if valStr == "" {
			return defaultVal
    }

    val := strings.Split(valStr, sep)

    return val
}