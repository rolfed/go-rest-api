package user

import (
	"fmt"
	"strconv"

	"github.com/rolfed/go-rest-api/src/database"
)

type User struct {
	Connection database.Connection
}

func QueryUserById(userID string) {

	// Convert userID string to int
	id, err := strconv.Atoi(userID)
	if err == nil {
		fmt.Println(id)
	}
}

func queryAllUsers() {

}
