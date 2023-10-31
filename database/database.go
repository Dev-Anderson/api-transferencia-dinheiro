package database

import (
	"api-transferencia/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error load Env ", err.Error())
	}
	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.HostDB, env.PortDB, env.UserDB, env.PasswordDB, env.DBName)

	db, err := sql.Open("postgres", stringDB)
	if err != nil {
		panic(err)
	}

	return db, err
}
