package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/rolfed/go-rest-api/src/env"
	"github.com/rolfed/go-rest-api/src/database"
	"github.com/rolfed/go-rest-api/src/helloworld"
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

  db, err := database.OpenDB()
		if err != nil {
			log.Fatal(err)
		} 
	// env := &database.Env{DB: db}
	defer db.Close()

	router.Route("/v1/api", func(r chi.Router) {
		// Routes
		// r.Mount("/user", u.Routes())
		r.Mount("/helloworld", helloworld.Resource{}.Routes())
	})

	return router
}

func main() {
	SERVER_PORT := env.Load("SERVER_PORT")
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


			fmt.Printf("\n ==> Starting server on port localhost:%s <==\n", SERVER_PORT)
			log.Fatal(http.ListenAndServe(":"+SERVER_PORT, router))
		}
