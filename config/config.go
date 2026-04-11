package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string
	AppEnv        string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
	SessionSecret string
}

var App Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	App = Config{
		AppPort:       getEnv("APP_PORT", "8080"),
		AppEnv:        getEnv("APP_ENV", "development"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBName:        getEnv("DB_NAME", "clientcompany"),
		JWTSecret:     getEnv("JWT_SECRET", "secret"),
		SessionSecret: getEnv("SESSION_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
