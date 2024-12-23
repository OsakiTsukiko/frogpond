package server

import (
	"log"

	"github.com/OsakiTsukiko/frogpond/server/api"
	"github.com/OsakiTsukiko/frogpond/server/handlers"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main program entry point
func Run() {
	router := gin.Default(func(engine *gin.Engine) {
		// do nothing for now
	})

	// AUTH (BAKED)

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	auth_group := router.Group("/auth", handlers.AuthMiddleware)
	auth_group.GET("/register", handlers.RegisterGET)
	auth_group.POST("/register", handlers.RegisterPOST)
	auth_group.GET("/login", handlers.LoginGET)
	auth_group.POST("/login", handlers.LoginPOST)

	req_auth_group := router.Group("/", handlers.ReqAuthMiddleware)
	req_auth_group.GET("/", handlers.HomeGET)
	req_auth_group.GET("/auth/app", handlers.AppGET)
	req_auth_group.POST("/auth/app", handlers.AppPOST)
	req_auth_group.POST("/auth/removeTokens", handlers.RemoveTokens)
	req_auth_group.POST("/auth/delete", handlers.DeleteUser)

	router.GET("/auth/logout", handlers.LogoutGET)

	// API

	api_req_auth := router.Group("/api", api.ReqAuthToken)
	api_req_auth.GET("/profile", api.GetProfile)

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
