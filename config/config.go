package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

type AppConfig struct {
	DB          DBConfig
	ENVIRONMENT string
	// Add other config sections here (e.g., Server, Auth)
}

var Config AppConfig

func Load() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Application configuration
	Config.ENVIRONMENT = os.Getenv("ENVIRONMENT")

	// Database configuration
	Config.DB = DBConfig{
		DSN: os.Getenv("DB"),
	}

}
