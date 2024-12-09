package handlers

import (
	"net/http"

	"github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// Redirect if Already Authentificated
func AuthMiddleware(c *gin.Context) {
	_, _, ok := UserFromSession(c)
	if ok /* is authentificated */ {
		c.Redirect(http.StatusFound, singleton.CFG.Server.DefaultRedirect)
		return
	}

	c.Next()
}
