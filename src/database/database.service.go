package database

import (
	"database/sql"
)

type Connection interface {
}

func NewConnection(db *sql.DB) Connection {
	return &connection{db}
}

type connection struct {
	db *sql.DB
} 
