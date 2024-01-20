package main

import (
	"fmt"
	"whois-api/configer"
	"whois-api/libs/logger"
	"whois-api/router"

	"github.com/gin-gonic/gin"
)

func webServer() {
	// inital configuer
	configer.InitialConfier()
	logger.InitialLogger()

	// set run mode
	if configer.Configer.AppMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// initial web server
	router.InitialRouter()

	router.Router.Run(fmt.Sprintf("0.0.0.0:%s", configer.Configer.Serve.Port))
}

func main() {
	webServer()
}
