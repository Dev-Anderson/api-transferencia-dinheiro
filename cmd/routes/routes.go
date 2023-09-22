package routes

import (
	"api-transferencia/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRoutes() {
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.GET("/user", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:cpf", controllers.DeleteUser)
	r.Run(":8080")
}
