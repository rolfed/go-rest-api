package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// User routes
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUser)
	router.Get("/", GetUsers)
	router.Delete("/{userID}", DeleteUser)
	router.Post("/", CreateUser)
	router.Put("/{userId}", UpdateUser)
	return router
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user := User{
		userID:    userID,
		email:     "test@gmail.com",
		password:  "password",
		createdOn: "Feb 14 2020 22:16:24 GMT-0800 (Pacific Standard Time)",
		lastLogin: "Feb 14 2020 22:16:24 GMT-0800 (Pacific Standard Time)",
	}
	log.Printf("\nGet User by id:%s\n", userID)
	render.JSON(w, r, user) // A chi router helper for serializing and returning json
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
