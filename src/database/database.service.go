package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Env struct {
	DB *sql.DB
	log *log.Logger
}


func OpenDB(dataSourceName string) (*sql.DB, error) {
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

