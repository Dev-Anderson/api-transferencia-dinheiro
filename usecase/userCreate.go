package usecase

import (
	"api-transferencia/database"
	"api-transferencia/schemas"
)

func CreateUser(u schemas.UserCreate) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `
		insert into "user" (name, document, "dateCreate", balance, "dateAtualization") values ($1, $2, now(), $3, now());`

	_, err = db.Exec(query, u.Name, u.Document, u.Balance)
	if err != nil {
		return err
	}

	return nil
}
