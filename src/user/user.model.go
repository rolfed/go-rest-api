package user

// User model
type User struct {
	userID   string `json:"primary_key"`
	email    string `json:"type:varchar(100);unique_index;not null"`
	password string `json:"type:string;not null"`
}
