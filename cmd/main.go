package main

import (
	"api-transferencia/internal/usecase/config"
	"api-transferencia/pkg/database"
)

func main() {
	database.ConnectionDB()
	config.TableExtractUser()
	config.TableUser()
	config.TableBalance()
	// routes.HandlerRoutes()
}
