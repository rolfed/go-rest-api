package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	userID   string `gorm:"primary_key"`
	email    string `gorm:"type:varchar(100);unique_index;not null"`
	password string `gorm:"type:string;not null"`
}
