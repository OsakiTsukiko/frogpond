package handlers

import (
	"net/http"

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

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
		"email":    email,
	})
}
