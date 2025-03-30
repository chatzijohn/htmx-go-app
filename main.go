package main

import (
	"log"
	"my-app/config"
	"my-app/db"
)

func main() {
	config.Load()

	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to initialize DB:", err)
	}

}
