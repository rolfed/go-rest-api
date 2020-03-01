package user

// import (
// 	"net/http"
// 
// 	"github.com/go-chi/chi"
// )

// Routes for user
// func Routes() *chi.Mux {
// 	router := chi.NewRouter()
// 	router.Get("/{userID}", getUser)
// 	router.Get("/", getAllUsers)
// 	router.Delete("/{userID}", deleteUser)
// 	router.Post("/", createUser)
// 	router.Put("/{userId}", updateUser)
// 	return router
// }

// func getUser(w http.ResponseWriter, r *http.Request) {
	// userID := chi.URLParam(r, "userID")

	// Query Database for user
  // Service.queryUserById(*sql.DB, userID)

	// Handler errr
	// log.Printf("\nGet User by id %+v\n", user)
	// render.JSON(w, r, user) // A chi router helper for serializing and returning json
// }

// func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Service.queryAllUsers(*sql.DB)
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 
// }
// 
// func createUser(w http.ResponseWriter, r *http.Request) {
// 
// }
// 
// func updateUser(w http.ResponseWriter, r *http.Request) {
// 
// }
