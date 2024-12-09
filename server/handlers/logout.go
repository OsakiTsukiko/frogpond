package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	ClearSession(c)
	c.Redirect(http.StatusFound, "/auth/login")
}
