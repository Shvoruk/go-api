package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:55001/go?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
