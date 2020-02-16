package user

// User model
type User struct {
	userID   string `json:"userID"`
	email    string `json:"email"`
	password string `json:"password"`
}
