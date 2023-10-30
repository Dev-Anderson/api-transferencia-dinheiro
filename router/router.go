package router

import (
	"api-transferencia/config"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	port, err := config.LoadEnv()
	if err != nil {
		log.Panic("Error load env ", err.Error())
	}
	router.Run("0.0.0.0:" + port.PortHttp)
}
