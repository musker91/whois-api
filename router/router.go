package router

import (
	"path/filepath"
	"whois-api/controllers"
	"whois-api/middlewares"
	"whois-api/utils"

	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func InitialRouter() {
	Router = gin.New()

	Router.LoadHTMLGlob(filepath.Join(utils.GetRootPath(), "views", "/*"))

	// set middlewares
	Router.Use(gin.Recovery())
	Router.Use(middlewares.AllowCors())

	// page
	{
		Router.GET("/", controllers.SiteHome)
	}

	// api
	{
		Router.GET("/api", controllers.WhoisQuery)
		Router.POST("/api", controllers.WhoisQuery)
	}
}
