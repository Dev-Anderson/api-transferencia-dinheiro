package routes

import (
	"api-transferencia/controllers"

	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		balance := main.Group("balance")
		{
			balance.GET("/:id", controllers.GetBalanceID)
		}
		transfer := main.Group("transfer")
		{
			transfer.POST("/", controllers.Transfer)
		}
	}

	return router
}
