package helloworld

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/go-chi/chi"

	"github.com/rolfed/go-rest-api/src/database"
)

// Routes for hello world
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getHelloWorld)
	router.Get("/{helloWorldID}", getHelloWorldById)
	router.Post("/", postHelloWorld)

	return router
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDB()
	helloWorlds, err := readHelloWorld(db)
	if err != nil {
		// Internal Server Error
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helloWorlds)
	defer db.Close()
}


func getHelloWorldById(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDB()
	helloWorldId := chi.URLParam(r, "helloWorldID") 
	if helloWorldId == "" {
		log.Fatal("hello world id is undefined")
	}

	// Parse id string to int
	id, err := strconv.Atoi(helloWorldId)
	if err != nil {
		log.Fatal("hello world id error ", err)
	}

	helloWorld, err := readHelloWorldById(db, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	} else {
		json.NewEncoder(w).Encode(helloWorld)
	}

	defer db.Close()
}

func postHelloWorld(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDB()
	
	// Parse Request Body

	// Pass json to repository

	defer db.Close()
}


