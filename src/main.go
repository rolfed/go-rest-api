package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"

	userTest "github.com/rolfed/go-rest-api/src/user"
)

// DataSourceName
const (
	host     = "localhost"
	port     = 5432
	user     = "dev"
	password = "password"
	dbname   = "dev"
	sslmode  = "disable"
)

// Routes for service
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will
	// signal through ctx.Done() that the request has timed out and
	// further processing should be stopped
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/v1/api", func(r chi.Router) {
		// Routes
		r.Mount("/user", userTest.Routes())
	})

	return router
}

func main() {
	port := "8001"
	router := Routes()

	walkFunc := func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panic("Logging err: %s\n", err.Error())
	}

	// Connect Database
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		// This will not be a connection error, but a DSN parse error
		// or another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	fmt.Printf("\n ==> Starting server on port localhost:%s <==\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
