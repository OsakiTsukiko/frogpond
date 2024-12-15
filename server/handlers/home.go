package handlers

import (
	"net/http"

	"github.com/OsakiTsukiko/frogpond/server/database"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

// main handler (/)
func HomeGET(c *gin.Context) {
	username, email, ok := UserFromSession(c)
	if !ok {
		ClearSession(c)
		c.Redirect(http.StatusFound, "/auth/login")
		// redirect to login if session invalid
	}

	/* c.JSON(http.StatusOK, gin.H{
		"username": username,
		"email":    email,
	}) */

	user, err := database.GetUserByUsername(username, sgl.DATABASE)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Error loading USER from DATABASE!",
		})
		return
	}

	if user == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "User not in DATABASE!!?",
		})
		return
	}

	tokens, err := database.GetUserTokens(user.ID, sgl.DATABASE)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Error getting tokens for user!",
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
		"email":    email,
		"tokens":   tokens,
	})
}
