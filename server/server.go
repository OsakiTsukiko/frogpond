package server

import (
	"github.com/OsakiTsukiko/frogpond/server/handlers"
	"github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main program entry point
func Run() {
	router := gin.Default(func(engine *gin.Engine) {
		// do nothing for now
	})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	auth_group := router.Group("/auth", handlers.AuthMiddleware)
	auth_group.GET("/register", handlers.RegisterGET)

	// run gin
	router.Run(":" + singleton.CFG.Server.Port)
}
