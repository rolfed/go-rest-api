package helloworld

import (
	"context"
	"net/http"
	"fmt"
	"time"
	"encoding/json"

	"github.com/go-chi/chi"

	"github.com/rolfed/go-rest-api/src/database"
)

type Resource struct {
	env *database.Env
}

// Routes for hello world
func (rs Resource) Routes(env *database.Env) *chi.Mux {
	rs.env = env

	router := chi.NewRouter()
	router.Get("/", rs.getHelloWorld)
	router.Get("/{id}", rs.getHelloWorldById)

	return router
}

func (rs Resource) getHelloWorld(w http.ResponseWriter, r *http.Request) {
	// Is the database connection alive
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	err := rs.env.DB.PingContext(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database down: %v", err), http.StatusFailedDependency)
	}

	helloWorlds, err := readHelloWorld(rs.env.DB)
	if err != nil {
		// Internal Server Error
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// for _, helloWorld := range helloWorlds {
	// 	fmt.Fprintf(w, "%s %s \n", helloWorld.ID, helloWorld.Description)
	// }
	json.NewEncoder(w).Encode(helloWorlds)
}


func (rs Resource) getHelloWorldById(w http.ResponseWriter, r *http.Request) {
	// Is the database connection alive
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	err := rs.env.DB.PingContext(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database down: %v", err), http.StatusFailedDependency)
	}

	// TODO parse id from url and convert it to int
	// helloWorld, err := readHelloWorldById(rs.env.DB, id)
	// if err != nil {
	// 	return err
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"test":}`))

}

