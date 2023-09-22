package database

import (
	"api-transferencia/pkg/env"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {
	host := env.LoadEnvironment().DB_host
	port := env.LoadEnvironment().DB_port
	user := env.LoadEnvironment().DB_user
	password := env.LoadEnvironment().DB_password
	dbName := env.LoadEnvironment().DB_name

	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", stringDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db, nil
}
