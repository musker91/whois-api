package main

import (
	"whois-api/configer"
	"whois-api/libs/logger"
	"whois-api/models"
	"whois-api/router"

	"github.com/gin-gonic/gin"
)

func webServer() {
	// inital configuer
	configer.InitialConfier()
	logger.InitialLogger()
	models.InitialRedis()

	defer models.RedisClient.Close()

	// initial web server
	router.InitialRouter()
	if configer.Configer.AppMode != "production" {
		gin.SetMode(gin.DebugMode)
	}

	router.Router.Run("0.0.0.0:8091")
}

func main() {
	webServer()
}
