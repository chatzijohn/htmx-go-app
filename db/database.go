package db

import (
	"database/sql"

	"my-app/config"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	var err error

	db, err := sql.Open("sqlite3", config.Config.DB.DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	// Verify connection and run migrations
	if err = db.Ping(); err != nil {
		return err
	}

	return Migrate(db)
}
