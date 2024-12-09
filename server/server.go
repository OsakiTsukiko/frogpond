package server

import (
	"github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main program entry point
func Run() {
	router := gin.Default(func(engine *gin.Engine) {
		// do nothing for now
	})

	// run gin
	router.Run(":" + singleton.CFG.Server.Port)
}
