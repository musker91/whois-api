package main

import (
	"whois-api/configer"
	"whois-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// inital configuer
	configer.InitialConfier()

	// initial web server
	router.InitialRouter()
	if configer.Configer.AppMode != "production" {
		gin.SetMode(gin.DebugMode)
	}

	router.Router.Run("0.0.0.0:8091")
}
