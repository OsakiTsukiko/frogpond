package handlers

import (
	"net/http"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main handler (/)
func HomeGET(c *gin.Context) {
	user_any, ok := c.Get("user")
	if !ok {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unable to ACCESS USER SESSION!!?",
		})
		return
	}
	user := user_any.(*d.User)

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
