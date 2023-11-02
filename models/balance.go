package models

import (
	"api-transferencia/database"
	"api-transferencia/schemas"
	"fmt"
)

func GetBalanceID(id int) []schemas.UserBalance {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `select id, balance from "user" where id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("Erro in get user balace", err)
	}
	defer rows.Close()

	var balance []schemas.UserBalance
	for rows.Next() {
		var b schemas.UserBalance
		if err := rows.Scan(&b.ID, &b.Balance); err != nil {
			fmt.Print(err)
		}
		balance = append(balance, b)
	}

	if err := rows.Err(); err != nil {
		fmt.Print(err)
	}

	return balance
}
