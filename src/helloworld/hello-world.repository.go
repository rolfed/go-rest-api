package helloworld

import (
	"database/sql"
	"fmt"
)

func readHelloWorld(db *sql.DB) ([]*HelloWorld, error) {
	sqlStatement := `SELECT * FROM hello_world_table;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	helloWorlds := make([]*HelloWorld, 0)
	for rows.Next() {
		helloWorld := new(HelloWorld)
		err := rows.Scan(&helloWorld.ID, &helloWorld.Description)
		if err != nil {
			return nil, err
		}

		helloWorlds = append(helloWorlds, helloWorld)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return helloWorlds, nil
}

// TODO add HelloWorld to return type in func
func readHelloWorldById(db *sql.DB, helloWorldId int) (HelloWorld, error) {
	sqlStatement := `SELECT * FROM hello_world_table WHERE id=$1;`
	res := HelloWorld{}
	row := db.QueryRow(sqlStatement, helloWorldId)

	var id int
	var description string
	switch err := row.Scan(&id, &description); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			res.ID = id
			res.Description = description
			// fmt.Println(db.QueryRowContext)
		default:
			panic(err)
	}

	return res, nil
}
