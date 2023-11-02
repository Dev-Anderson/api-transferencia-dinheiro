package routes

import (
	"api-transferencia/config"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	configRoutes(router)

	port := config.LoadEnv().PortHttp
	router.Run("localhost:" + port)
}
