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
func readHelloWorldById(db *sql.DB, id int) (error) {
	// TODO fix code below
	// if (id == 0) {
	// 	return fmt.Errorf("Repository ID not correctly set"), nil
	// }

	// helloWorld := new(HelloWorld)
	helloWorldId := id;

	sqlStatement := `SELECT id FROM hello_world_table WHERE id=$l`
	row := db.QueryRow(sqlStatement, helloWorldId)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Println(db.QueryRowContext)
	default:
		panic(err)
	}

	return nil
}
