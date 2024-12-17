package handlers

import (
	"net/http"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func RemoveTokens(c *gin.Context) {
	user_any, ok := c.Get("user")
	if !ok {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unable to ACCESS USER SESSION!!?",
		})
		return
	}
	user := user_any.(*d.User)

	user.RemoveAllTokens(sgl.DATABASE)
	c.Redirect(http.StatusFound, "/")
}
