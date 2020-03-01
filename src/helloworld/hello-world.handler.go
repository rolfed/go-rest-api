package helloworld

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/go-chi/chi"

	"github.com/rolfed/go-rest-api/src/database"
)

type Resource struct {
	env *database.Env
	dsn string
}

// Routes for hello world
func (rs Resource) Routes(env *database.Env, dsn string) *chi.Mux {
	rs.env = env
	rs.dsn = dsn

	router := chi.NewRouter()
	router.Get("/", rs.getHelloWorld)
	router.Get("/{helloWorldID}", rs.getHelloWorldById)

	return router
}

func (rs Resource) getHelloWorld(w http.ResponseWriter, r *http.Request) {
	helloWorlds, err := readHelloWorld(rs.env.DB)
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
	helloWorldId := chi.URLParam(r, "helloWorldID") 
	if helloWorldId == "" {
		log.Fatal("hello world id is undefined")
	}

	// Parse id string to int
	id, err := strconv.Atoi(helloWorldId)
	if err != nil {
		log.Fatal("hello world id error ", err)
	}

	helloWorld, err := readHelloWorldById(rs.env.DB, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	} else {
		json.NewEncoder(w).Encode(helloWorld)
	}
}

