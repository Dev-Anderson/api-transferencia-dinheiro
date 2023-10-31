package main

import (
	"api-transferencia/usecase"
	"fmt"
)

func main() {
	fmt.Println("API-tranferencia")
	usecase.TableUser()
	usecase.TableBalance()
	usecase.FKBalanceIDUserOrigin()
	usecase.FKBalanceIDUserDestiny()
	fmt.Println("")

	// user := schemas.UserCreate{
	// 	Name:     "Anderson",
	// 	Document: "1234",
	// 	Balance:  100.0,
	// }
	// err := usecase.CreateUser(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// router.Initialize()
}
