package user

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type userRepository struct {
	db *sql.DB
}

type Container interface {
	db *sql.DB

	userRepository *userRepository
}

func (c *Container) queryServiceById(userID string) {
	c.userRepository = &userRepository{db}

	// Convert userID string to int
	id, err := strconv.Atoi(userID)
	if err == nil {
		fmt.Println(id)
	}
}

func (c *Container) queryAllServices() {
	db := c.db

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
