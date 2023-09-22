package env

import (
	"api-transferencia/internal/types"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() types.Environment {
	godotenv.Load()
	return types.Environment{
		DB_host:     os.Getenv("db_host"),
		DB_port:     os.Getenv("db_port"),
		DB_user:     os.Getenv("db_user"),
		DB_password: os.Getenv("db_password"),
		DB_name:     os.Getenv("db_name"),
	}
}
