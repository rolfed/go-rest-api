package user

type User struct {
	userID    string `json:"userID"`
	email     string `json:"email"`
	password  string `json:"password"`
	createdOn string `json:"createdOn"`
	lastLogin string `json:"lastLogin"`
}
