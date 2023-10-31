package handler

import (
	"api-transferencia/schemas"
	"api-transferencia/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user, err := usecase.GetUser()
	if err != nil {
		log.Fatal("Error when fetching data from the database ", err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var user schemas.UserCreate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usecase.CreateUser(user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario invalido"})
		return
	}

	err = usecase.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Nao foi possivel deletar o usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario excluido com sucesso",
	})

}

func GetUserID(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuario invalido"})
		return
	}

	user, err := usecase.GetUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}
