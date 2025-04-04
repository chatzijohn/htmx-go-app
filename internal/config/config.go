package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

type ServerConfig struct {
	PORT string
	HOST string
}

type AppConfig struct {
	DB          DBConfig
	ENVIRONMENT string
	SERVER      ServerConfig
	// Add other config sections here (e.g., Server, Auth)
}

func Load() *AppConfig {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Application configuration
	environment := os.Getenv("ENVIRONMENT")

	// Server configuration
	server := ServerConfig{
		HOST: getEnvWithDefault("HOST", "127.0.0.1"),
		PORT: getEnvWithDefault("PORT", "8080"),
	}

	// Database configuration
	db := DBConfig{
		DSN: os.Getenv("DB"),
	}

	return &AppConfig{ENVIRONMENT: environment, DB: db, SERVER: server}

}

// Helper function to get env with fallback
func getEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
