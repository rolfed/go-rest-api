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
	psqlConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Opening a driver typically will not attempt to connect
	// to the database
	pool, err := sql.Open("postgres", psqlConnection)
	if err != nil {
		// This will not be connection error, but a DSN
		// parse error
		log.Fatal("unable to use data source name", err)
	}

	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

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

// Ping the database to verify DSN provided by the user is valid and
// the server accessible. If the pin fails exit the program with an error
// func Ping(ctx context.Context) {
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()
//
// 	if err := pool.PingContext(ctx); err != nil {
// 		log.Fatal("unable to connect to database: %v", err)
// 	}
// }
