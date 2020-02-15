package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/rolfed/go-rest-api/src/config"
	"github.com/rolfed/go-rest-api/src/user"
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
		r.Mount("/user", user.Routes())
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
	config.InitDBConnection()

	fmt.Printf("Starting server on port localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
