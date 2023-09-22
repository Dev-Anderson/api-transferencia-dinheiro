package main

import (
	"api-transferencia/cmd/routes"
	"api-transferencia/pkg/database"
	"fmt"
)

func main() {
	fmt.Println("Conexao com o banco de dados")
	database.ConnectionDB()
	routes.HandlerRoutes()
}
