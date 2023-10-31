package usecase

import (
	"api-transferencia/database"
	"api-transferencia/schemas"
	"fmt"
)

func GetUser() ([]schemas.User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
	select 
		id, name, document, "dateCreate", balance, "dateAtualization"
	from "user"
	order by id`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro in Get user ", err.Error())
	}
	defer rows.Close()

	var users []schemas.User
	for rows.Next() {
		var user schemas.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Document, &user.DateCreate, &user.Balance, &user.DateAtualization); err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return users, nil
}
