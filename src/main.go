package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/callistaenterprise/goblog/common/tracing"
	"github.com/callistaenterprise/goblog/dataservice/dbclient"
	"github.com/callistaenterprise/goblog/dataservice/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/rolfed/go-rest-api/src/config"
	"github.com/spf13/viper"
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

var DB config.IGormClient
var appName "go-rest-api"

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
	config.DBClient = &dbclient.GormClient{}
	config.DBClient.SetupDB()
	initializeTracing()

	handleSigterm(func() {
		logrus.Infoln("Captured Ctrl+C")
		service.DBClient.Close()
	})

	fmt.Printf("\n ==> Starting server on port localhost:%s <==\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

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

	router.Route("/v1", func(r chi.Router) {
		// Routes
		r.Mount("/api/user", user.Routes())
	})

	return router
}

func initializeTracing() {
	tracing.InitTracing(viper.GetString("zipkin_server_url"), appName)
}
