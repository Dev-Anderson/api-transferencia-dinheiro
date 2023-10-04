package main

import (
	"api-transferencia/cmd/routes"
	"api-transferencia/internal/usecase/config"
	"api-transferencia/pkg/database"
	"fmt"
)

func main() {
	database.ConnectionDB()
	fmt.Println("Created tables")
	config.TableExtractUser()
	config.TableUser()
	config.TableBalance()
	fmt.Println("Finished creating tables")
	fmt.Println("")
	fmt.Println("Starting creation of FK")
	config.FKExtractUserIDOrigem()
	config.FKExtractUserIDDestiny()
	config.FKBalanceIDUser()
	fmt.Println("Finished creating FK")
	fmt.Println("")
	fmt.Println("Starting the routes")
	routes.HandlerRoutes()
}
