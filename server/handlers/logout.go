package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutGET(c *gin.Context) {
	ClearSession(c)
	c.Redirect(http.StatusFound, "/auth/login")
}
