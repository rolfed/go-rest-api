package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
)

// DNS
const (
	host     = "localhost"
	port     = 5432
	user     = "dev"
	password = "password"
	dbname   = "dev"
	sslmode  = "disable"
)

var pool *sql.DB // Database connection pool

// Create PostGrest database connection
func ConnectDB() {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	// Opening a driver typically will not attempt to connect
	// to the database
	pool, err := sql.Open("postgres", dns)
	if err != nil {
		// This will not be connection error, but a DSN
		// parse error
		log.Fatal("unable to use data source name", err)
	}

	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	defer func() {
		signal.Stop(appSignal)
		cancel()
	}()

	go func() {
		select {
		case <-appSignal:
			cancel()
		case <-ctx.Done():
		}

	}()

	err = pool.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected to database")
}
