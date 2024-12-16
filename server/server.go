package server

import (
	"log"

	"github.com/OsakiTsukiko/frogpond/server/handlers"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main program entry point
func Run() {
	router := gin.Default(func(engine *gin.Engine) {
		// do nothing for now
	})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.GET("/", handlers.HomeGET)

	auth_group := router.Group("/auth", handlers.AuthMiddleware)
	auth_group.GET("/register", handlers.RegisterGET)
	auth_group.POST("/register", handlers.RegisterPOST)
	auth_group.GET("/login", handlers.LoginGET)
	auth_group.POST("/login", handlers.LoginPOST)

	req_auth_group := router.Group("/auth", handlers.ReqAuthMiddleware)
	req_auth_group.GET("/app", handlers.AppGET)
	req_auth_group.POST("/app", handlers.AppPOST)
	req_auth_group.POST("/removeTokens", handlers.RemoveTokens)
	req_auth_group.POST("/delete", handlers.DeleteUser)

	router.GET("/auth/logout", handlers.LogoutGET)

	if sgl.CFG.Server.UseHTTPS {
		err := router.RunTLS(":"+sgl.CFG.Server.Port, sgl.CFG.Server.FullChain, sgl.CFG.Server.PrivKey)
		if err != nil {
			log.Fatalf("🚩 Failed to start HTTPS server: %v", err)
		}
	} else {
		// run gin
		err := router.Run(":" + sgl.CFG.Server.Port)
		if err != nil {
			log.Fatalf("🚩 Failed to start HTTP server: %v", err)
		}
	}
}
