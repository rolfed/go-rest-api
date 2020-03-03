package helloworld

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"
	"fmt"

	"github.com/go-chi/chi"

	"github.com/rolfed/go-rest-api/src/database"
)

type Resource struct {}

// Routes for hello world
func (rs Resource) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", rs.getHelloWorld)
	router.Get("/{helloWorldID}", rs.getHelloWorldById)
	router.Post("/", rs.postHelloWorld)
	router.Put("/{helloWorldID}", rs.putHelloWorld)

	return router
}

func (rs Resource) getHelloWorld(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	helloWorlds, err := readHelloWorld(db)
	if err != nil {
		// Internal Server Error
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helloWorlds)
}


func (rs Resource) getHelloWorldById(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	helloWorldId := chi.URLParam(r, "helloWorldID") 
	if helloWorldId == "" {
		log.Fatal("hello world id is undefined")
	}

	// Parse id string to int
	id, err := strconv.Atoi(helloWorldId)
	if err != nil {
		log.Printf("hello world id %s is not valid", err)
	}

	helloWorld, err := readHelloWorldById(db, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return 
	} 

  json.NewEncoder(w).Encode(helloWorld)

}

func (rs Resource) postHelloWorld(w http.ResponseWriter, r *http.Request) {
	var helloWorld HelloWorld

	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	// Decode Request Body
	helloWorldRequest := json.NewDecoder(r.Body).Decode(&helloWorld)
	if helloWorldRequest != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "HW: %+v", helloWorld)

	// Pass data to repository
	err = createHelloWorld(db, helloWorld)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (rs Resource) putHelloWorld(w http.ResponseWriter, r *http.Request) {
	var helloWorld HelloWorld

	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	helloWorldId := chi.URLParam(r, "helloWorldID") 
	if helloWorldId != "" {
		log.Printf("Hello world with id: %s ", helloWorldId)
	}

	// Parse id string to int
	id, err := strconv.Atoi(helloWorldId)
	if err != nil {
		log.Printf("hello world id %s is not valid", err)
	}

	// Decode Request Body
	helloWorldRequest := json.NewDecoder(r.Body).Decode(&helloWorld)
	if helloWorldRequest != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateHelloWorld(db, id, helloWorld) 
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	} 

	w.WriteHeader(204)
}
