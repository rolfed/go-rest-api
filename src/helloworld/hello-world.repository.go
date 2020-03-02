package helloworld

import (
	"database/sql"
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

func readHelloWorldById(db *sql.DB, helloWorldId int) (HelloWorld, error) {
	sqlStatement := `SELECT * FROM hello_world_table WHERE id=$1;`
	res := HelloWorld{}
	row := db.QueryRow(sqlStatement, helloWorldId)

	var id int
	var description string
	switch err := row.Scan(&id, &description); err {
		case sql.ErrNoRows:
			return res, err 
		case nil:
			res.ID = id
			res.Description = description
		default:
			panic(err)
	}

	return res, nil
}

func createHelloWorld(db *sql.DB, helloWorld HelloWorld) (error) {
	sqlStatement := `
		INSERT INTO hello_world_table(id, description) 
		VALUES($1, $2);`
	
		_, err := db.Exec(sqlStatement, helloWorld.ID, helloWorld.Description)
		if err != nil {
		  return err	
		}

	return nil
}
