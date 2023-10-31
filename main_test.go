package main

import (
	"api-transferencia/handler"
	"api-transferencia/schemas"
	"api-transferencia/usecase"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var totalTest, testsPassed, testsFailed int

func TestMain(m *testing.M) {
	code := m.Run()

	println("Total de testes: ", totalTest)
	println("Testes passados: ", testsPassed)
	println("Testes falhados: ", testsFailed)

	if code == 0 {
		println("Todos os testes passaram!")
	} else {
		println("Alguns testes falharam!")
	}

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	totalTest++
	user := schemas.UserCreate{
		Name:     "Anderson",
		Document: "123456789",
		Balance:  100.0,
	}

	err := usecase.CreateUser(user)

	if err != nil {
		t.Errorf("Erro ao criar usuario: %v", err)
		testsFailed++
	}
	testsPassed++
}

func TestGetUser(t *testing.T) {
	totalTest++
	users, err := usecase.GetUser()
	if err != nil {
		t.Fatalf("Erro ao buscar os usuarios %v", err)
		testsFailed++
	}

	if len(users) == 0 {
		t.Fatal("Nenhum usuario retornado")
		testsFailed++
		return
	}

	testsPassed++

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

func TestGetUserHandler(t *testing.T) {
	totalTest++
	//criando um objeto de tipo gin.context, para user nos testes
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	handler.GetUser(c)

	//verificando a resposta HTTP e o corpo da resposta
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	if c.Writer.Status() == http.StatusOK {
		testsPassed++
	} else {
		testsFailed++
	}
}

func TestPostUserHandler(t *testing.T) {
	totalTest++
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/api/v1/user", handler.PostUser)

	requestBody := `{
		"name": "Anderson da silva", 
		"document": "1234", 
		"balance": 10.15
	}`
	w := performRequest(r, "POST", "/api/v1/user", requestBody)
	assert.Equal(t, http.StatusOK, w.Code)

	if w.Code == http.StatusOK {
		testsPassed++
	} else {
		testsFailed++
	}

}

func performRequest(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
