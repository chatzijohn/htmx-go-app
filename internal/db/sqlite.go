package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"my-app/internal/config"
	"my-app/pkg/data"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	cfg := config.Load() // TODO: PASS THE CFG FROM SERVER INSTEAD OF HERE
	db, err := sql.Open("sqlite3", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	// Verify connection and run migrations
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Perform migration (this will call Migrate, which will execute the migration scripts)
	if err := Migrate(db, cfg); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB, cfg *config.AppConfig) error {
	// Read migrations file
	migrationSQL, err := os.ReadFile("internal/db/migrations.sql")
	if err != nil {
		return fmt.Errorf("could not read migrations file: %v", err)
	}

	// Prepare and execute the migration SQL
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		return fmt.Errorf("error executing migrations: %v", err)
	}

	log.Println("Database migration completed")

	// If in development environment, run seed
	if strings.ToLower(cfg.ENVIRONMENT) == "development" {
		return Seed(db)
	}
	return nil
}

func Seed(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	stmt, err := tx.Prepare(`
		INSERT OR IGNORE INTO countries (id, name, code, capital, continent)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing insert statement: %v", err)
	}
	defer stmt.Close()

	for _, country := range data.Countries {
		// Generate UUID
		country.ID = uuid.New()

		_, err := stmt.Exec(country.ID, country.Name, country.Code, country.Capital, country.Continent)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting country %s: %v", country.Name, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	log.Println("Database seeding completed succesfully.")
	return nil
}
