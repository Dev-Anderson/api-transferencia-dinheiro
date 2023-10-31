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
		user := main.Group("user")
		{
			user.GET("/", handler.GetUser)
			user.POST("/", handler.PostUser)
			user.DELETE("/:id", handler.DeleteUser)
			user.GET("/:id", handler.GetUserID)
		}
	}

	return router
}
