package handlers

import (
	"net/http"
	"net/url"

	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// Redirect if Already Authentificated
func AuthMiddleware(c *gin.Context) {
	user, ok := UserFromSession(c, sgl.DATABASE)
	if user != nil && ok /* is authentificated */ {
		c.Redirect(http.StatusFound, sgl.CFG.Server.DefaultRedirect)
		return
	}

	c.Next()
}

func ReqAuthMiddleware(c *gin.Context) {
	user, ok := UserFromSession(c, sgl.DATABASE)
	if user == nil || !ok /* is not authentificated */ {
		// TODO: check if the following is safe
		c.Redirect(http.StatusFound, "/auth/login?redirect="+url.QueryEscape(c.Request.URL.String()))
		return
	}

	c.Set("user", user)
	c.Next()
}
