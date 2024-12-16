package handlers

import (
	"net/http"

	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	user, ok := UserFromSession(c, sgl.DATABASE)
	if user == nil || !ok {
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	_ = user.Delete(sgl.DATABASE)
	c.Redirect(http.StatusFound, "/auth/login")
	return
}
