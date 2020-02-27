package user

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type service struct {
	db *sql.DB
}

type Service interface {
	queryAllUsers()
	queryUserById(userID string)
}

func (user *service) queryServiceById(userID string) {
	// Convert userID string to int
	id, err := strconv.Atoi(userID)
	if err == nil {
		fmt.Println(id)
	}
}

func (user *service) queryAllServices() {
	db := user.db

	var id, email, password string

	rows, err := db.Query("SELECT id, email, password FROM user_table")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &email, &password)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(id, email, password)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

}
