package main

import (
	"api-transferencia/schemas"
	"api-transferencia/usecase"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := schemas.UserCreate{
		Name:     "Anderson",
		Document: "123456789",
		Balance:  100.0,
	}

	err := usecase.CreateUser(user)

	if err != nil {
		t.Errorf("Erro ao criar usuario: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	users, err := usecase.GetUser()
	if err != nil {
		t.Fatalf("Erro ao buscar os usuarios %v", err)
	}

	if len(users) == 0 {
		t.Fatal("Nenhum usuario retornado")
		return
	}

	//caso queira exibir o resultado da consulta
	// for i, user := range users {
	// 	t.Logf("Usuario %d:", i+1)
	// 	t.Logf("ID: %d", user.ID)
	// 	t.Logf("Name: %s", user.Name)
	// 	t.Logf("Document: %s", user.Document)
	// 	t.Logf("Data criacao: %s", user.DateCreate)
	// 	t.Logf("Balance: %f", user.Balance)
	// 	t.Logf("Data Atualizacao: %s", user.DateAtualization)

	// }

}
