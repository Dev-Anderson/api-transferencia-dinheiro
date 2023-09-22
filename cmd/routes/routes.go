package routes

import (
	"api-transferencia/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRoutes() {
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.Run(":8080")
}
