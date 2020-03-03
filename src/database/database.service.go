package database

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	"github.com/rolfed/go-rest-api/src/env"
)

// DataSourceName

type Env struct {
	DSN string
	DB *sql.DB
	log *log.Logger
}


func OpenDB() (*sql.DB, error) {
	DB_HOST := env.Load("DB_HOST")
	DB_PORT := env.Load("DB_PORT")
	DB_USER := env.Load("DB_USER")
	DB_PASSWORD := env.Load("DB_PASSWORD")
	DB_NAME := env.Load("DB_NAME")
	DB_SSL_MODE := env.Load("DB_SSL_MODE")

	// Connect Database
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL_MODE)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		// This will not be a connection error, but a DSN parse error
		// or another initialization error.
		panic(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(48)
	db.SetMaxOpenConns(48)

	return db, nil
}

