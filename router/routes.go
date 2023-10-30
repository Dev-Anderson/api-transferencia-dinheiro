package router

import (
	"api-transferencia/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", handler.Home)
		}
	}

	return router
}
