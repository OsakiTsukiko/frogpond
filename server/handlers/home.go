package handlers

import (
	"net/http"

	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main handler (/)
func HomeGET(c *gin.Context) {
	user, ok := UserFromSession(c, sgl.DATABASE)
	if !ok {
		ClearSession(c)
		c.Redirect(http.StatusFound, "/auth/login")
		// redirect to login if session invalid
	}

	if user == nil { // this check might be useless? // TODO: CHECK IF THIS IS USELESS
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "User not in DATABASE!!?",
		})
		return
	}

	tokens, err := user.GetTokens(sgl.DATABASE)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Error getting tokens for user!",
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": user.Username,
		"email":    user.Email,
		"tokens":   tokens,
	})
}
