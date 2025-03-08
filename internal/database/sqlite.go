package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(connectionData string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connectionData)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA busy_timeout = 5000;")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database initialized with WAL mode, shared cache, and foreign keys enabled")
	return db, nil
}
