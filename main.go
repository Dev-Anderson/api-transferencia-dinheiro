package main

import (
	"api-transferencia/routes"
	"api-transferencia/usecase"
)

func main() {
	//create table
	usecase.TableBalance()
	usecase.TableUser()
	usecase.FKBalanceIDUserDestiny()
	usecase.FKBalanceIDUserOrigin()
	routes.Initialize()
}
