package controllers

import (
	"api-transferencia/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	user, err := models.GetUser()
	if err != nil {
		fmt.Println("Validacao da consulta do user", err.Error())
	}

	c.JSON(http.StatusOK, user)
}
