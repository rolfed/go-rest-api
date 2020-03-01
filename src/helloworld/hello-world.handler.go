package helloworld

import (
	"context"
	"net/http"
	"database/sql"
	"fmt"
	"time"
	"github.com/go-chi/chi"

	"github.com/rolfed/go-rest-api/src/database"
)

type HelloWorld struct {
	ID  string 
	Description string
}

type Resource struct {
	env *database.Env
}

// Routes for hello world
func (rs Resource) Routes(env *database.Env) *chi.Mux {
	rs.env = env

	router := chi.NewRouter()
	router.Get("/", rs.getHelloWorld)

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


	helloWorlds, err := allHelloWorld(rs.env.DB)
	if err != nil {
		// Internal Server Error
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, helloWorld := range helloWorlds {
		fmt.Fprintf(w, "%s %s", helloWorld.ID, helloWorld.Description)
	}

}

func allHelloWorld(db *sql.DB) ([]*HelloWorld, error) {
	rows, err := db.Query(`SELECT * FROM hello_world_table`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	helloWorlds := make([]*HelloWorld, 0)
	for rows.Next() {
		helloWorld := new(HelloWorld)
		err := rows.Scan(&helloWorld.ID, &helloWorld.Description)
		if err != nil {
			return nil, err
		}

		helloWorlds = append(helloWorlds, helloWorld)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return helloWorlds, nil
}
