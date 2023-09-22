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

func CreateUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created sucess"})
}

func DeleteUser(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	if err := models.DeleteUser(cpf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensage": "user successfully deleted"})
}
