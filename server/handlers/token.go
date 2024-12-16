package handlers

import (
	"net/http"

	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func RemoveTokens(c *gin.Context) {
	user, ok := UserFromSession(c, sgl.DATABASE)
	if !ok {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unable to load USER from SESSION!!?",
		})
		return
	}

	if user == nil { // might be useless // TODO: CHECK IF THIS IS USELESS
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "USER not in DATABASE?!",
		})
		return
	}

	user.RemoveAllTokens(sgl.DATABASE)
	c.Redirect(http.StatusFound, "/")
}
