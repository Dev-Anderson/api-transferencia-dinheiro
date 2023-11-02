package controllers

import (
	"api-transferencia/models"
	"api-transferencia/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	var transfer schemas.Transfer
	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.TransferBalance(transfer.IDOrigin, transfer.IDDestiny, transfer.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transferencia realizada com sucesso"})
}
