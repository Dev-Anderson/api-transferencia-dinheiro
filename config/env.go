package config

import (
	"api-transferencia/schemas"
	"bufio"
	"log"
	"os"
	"strings"
)

func setEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func LoadEnv() (schemas.ConfigEnv, error) {
	err := setEnv(".env")
	if err != nil {
		log.Fatal("Error set env ", err)
	}
	return schemas.ConfigEnv{
		HostDB:     os.Getenv("API-TRANSFERENCIA-HOSTDB"),
		PortDB:     os.Getenv("API-TRANSFERENCIA-PORTDB"),
		UserDB:     os.Getenv("API-TRANSFERENCIA-USERDB"),
		PasswordDB: os.Getenv("API-TRANSFERENCIA-PASSWORDDB"),
		DBName:     os.Getenv("API-TRANSFERENCIA-DBNAME"),
	}, nil
}
