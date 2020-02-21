package user

import (
	"fmt"
	"strconv"
)

func queryUserById(userID string) {

	// Convert userID string to int
	id, err := strconv.Atoi(userID)
	if err == nil {
		fmt.Println(id)
	}
}

func queryAllUsers() {

}
