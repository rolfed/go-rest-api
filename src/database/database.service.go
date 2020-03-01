package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Env struct {
	DB *sql.DB
}


func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		// This will not be a connection error, but a DSN parse error
		// or another initialization error.
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db, nil
}
