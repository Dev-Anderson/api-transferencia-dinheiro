package models

import (
	"api-transferencia/pkg/database"
	"fmt"
)

type User struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Cpf          string  `json:"cpf"`
	BalanceStart float64 `json:"balance_start"`
	CreateAt     string  `json:"create_at"`
}

func GetUser() ([]User, error) {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Falha na conexao com o banco de dados", err.Error())
	}
	defer db.Close()

	query := "select id, name, cpf, balance_start, create_at from user_balance order by id"
	var users []User

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Falha ao rodar a consulta", err.Error())
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Cpf, &user.BalanceStart, &user.CreateAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return users, nil

}

func CreateUser(u User) error {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO user_balance (name, cpf, balance_start, create_at) VALUES($1, $2, $3, now())")

	_, err2 := db.Exec(query, u.Name, u.Cpf, u.BalanceStart)
	if err2 != nil {
		return err
	}

	return nil
}
