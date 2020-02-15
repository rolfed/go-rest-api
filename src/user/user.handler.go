package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes for user
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUser)
	router.Get("/", GetUsers)
	router.Delete("/{userID}", DeleteUser)
	router.Post("/", CreateUser)
	router.Put("/{userId}", UpdateUser)
	return router
}

// GetUser by id
// response User struct
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

// GetUsers
// response Users array
func GetUsers(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
