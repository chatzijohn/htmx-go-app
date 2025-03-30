package db

import (
	"database/sql"
	"log"
)

// Migrate runs all database migrations
func Migrate(db *sql.DB) error {
	// Countries table
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS countries (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			code TEXT NOT NULL UNIQUE,
			capital TEXT NOT NULL UNIQUE,
			continent TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Database migration completed")
	return nil
}
