package main

import (
	"api-transferencia/cmd/routes"
	"api-transferencia/internal/usecase/config"
	"api-transferencia/pkg/database"
	"fmt"
)

func main() {
	fmt.Println("Conexao com o banco de dados")
	database.ConnectionDB()
	config.TableUser()
	routes.HandlerRoutes()
}
