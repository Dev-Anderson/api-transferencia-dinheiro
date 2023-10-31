package usecase

import (
	"api-transferencia/database"
	"log"
)

func UpdateUser(id int, name, document string, balance float64) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		update "user"
		set name = $1, document = $2, balance = $3, dateAtualization = now()
		where id = $4`

	_, err = db.Exec(query, name, document, balance, id)
	if err != nil {
		log.Fatalln("Erorr", err)
	}
}
