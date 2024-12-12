package handlers

import (
	"net/http"
	"net/url"

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

func ReqAuthMiddleware(c *gin.Context) {
	_, _, ok := UserFromSession(c)
	if !ok /* is authentificated */ {
		// TODO: check if the following is safe
		c.Redirect(http.StatusFound, "/auth/login?redirect="+url.QueryEscape(c.Request.URL.String()))
		return
	}

	c.Next()
}
