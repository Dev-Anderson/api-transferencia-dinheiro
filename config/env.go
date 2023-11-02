package config

import (
	"api-transferencia/schemas"
	"bufio"
	"log"
	"os"
	"strings"
)

func setEnv(fileName string) error {
	file, err := os.Open(fileName)
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

func LoadEnv() schemas.Env {
	err := setEnv(".env")
	if err != nil {
		log.Fatal("Error set env", err)
	}
	return schemas.Env{
		Host:     os.Getenv("api_transferencia_db_host"),
		Port:     os.Getenv("api_transferencia_db_port"),
		User:     os.Getenv("api_transferencia_db_user"),
		Password: os.Getenv("api_transferencia_db_password"),
		DBName:   os.Getenv("api_transferencia_db_name"),
		PortHttp: os.Getenv("api_transferencia_port_http"),
	}
}
