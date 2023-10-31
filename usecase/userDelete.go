package usecase

import "api-transferencia/database"

func DeleteUser(id int) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `DELETE FROM "user" WHERE ID = $1;`

	_, err = db.Exec(query, id)
	if err != nil {
		return nil
	}

	return nil
}
