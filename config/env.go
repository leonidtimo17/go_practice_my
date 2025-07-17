package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
	}
	log.Println(".env file loaded")
}

func getString(key, defaultValue string) string {
	value := os.Getenv("DATABASE_URL")
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv("DATABASE_URL")
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv("DATABASE_URL")
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

type DatabaseConfig struct {
	url string
}

// Default value

// Функции для получения int and bool

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}
