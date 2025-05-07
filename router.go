package main

import (
	"thunder_hoster/handler"
	"thunder_hoster/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handler.MainPage)
	router.POST("/auth", handler.PasswordAuthenticator)

	mapGroup := router.Group("/map")
	mapGroup.Use(middleware.AccessControlMiddleware())
	{
		mapGroup.GET("/", handler.SendFile)
		mapGroup.POST("/", handler.SendFile)
	}

	return router
}
