package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() *sql.DB {
	if db != nil {
		log.Println("DB already exist, returning that")
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "home_tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database initiated correctly")
	return db
}
