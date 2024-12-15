package handlers

import (
	"net/http"

	"github.com/OsakiTsukiko/frogpond/server/database"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func RemoveTokens(c *gin.Context) {
	username, _, ok := UserFromSession(c)
	if !ok {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unable to load USER from SESSION!!?",
		})
		return
	}
	user, err := database.GetUserByUsername(username, sgl.DATABASE)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unable to ACCESS user from DATABASE!",
		})
		return
	}

	if user == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "USER not in DATABASE?!",
		})
		return
	}

	database.RemoveAllTokensForUser(user.ID, sgl.DATABASE)
	c.Redirect(http.StatusFound, "/")
}
