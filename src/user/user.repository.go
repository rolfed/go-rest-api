package user

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/rolfed/go-rest-api/src/config"
)

func queryUserById(userID string) {

	// Convert userID string to int
	id, err := strconv.Atoi(userID)
	if err == nil {
		fmt.Println(id)
	}

	// TODO remove when done
	row := config.DB.QueryRow(`SELECT * FROM user_table WHERE user_id=$1`)
	switch err := row.Scan(&userID); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Println(userID)
	default:
		panic(err)
	}

	// result := []User{}

	// for rows.Next() {
	// 	var user User
	// 	err = rows.Scan(&user.userID, &user.email, &user.password)
	// 	if err != nil {
	// 		log.Fatalf("Scan: %v", err)
	// 	}
	// 	result = append(result, user)

	// }
}
