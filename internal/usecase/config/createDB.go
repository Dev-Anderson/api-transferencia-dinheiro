package config

import (
	"api-transferencia/pkg/database"
	"database/sql"
	"fmt"
)

func verifyTable(db *sql.DB, nameTable string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)")
	var exists bool
	err := db.QueryRow(query, nameTable).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func TableUser() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connect database", err.Error())
	}
	defer db.Close()

	tableExists, err := verifyTable(db, "user")
	if tableExists == false {
		createTable := `
			CREATE TABLE user_balance (
				ID SERIAL PRIMARY KEY, 
				NAME VARCHAR(255), 
				CPF VARCHAR(20),
				BALANCE_START FLOAT,  
				CREATE_AT DATE
			)
		`
		_, err := db.Exec(createTable)
		if err != nil {
			fmt.Println("Erro ao criar a tabela", err.Error())
		}
		fmt.Println("Tabela criada com sucesso!")
	} else if tableExists != false {
		panic(err)
	} else {
		fmt.Println("A tabela ja existe")
	}
}
