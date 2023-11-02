package controllers

import (
	"api-transferencia/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBalanceID(c *gin.Context) {
	id := c.Params.ByName("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID usuario invalido"})
		return
	}

	userBalance := models.GetBalanceID(idUser)
	c.JSON(http.StatusOK, userBalance)
}
