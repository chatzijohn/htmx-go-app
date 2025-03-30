package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string // Data Source Name
}

type AppConfig struct {
	DB DBConfig
	// Add other config sections here (e.g., Server, Auth)
}

var Config AppConfig

func Load() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Database configuration
	Config.DB = DBConfig{
		DSN: os.Getenv("DB"),
	}

}
