package api

import (
	"net/http"
	"strings"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func ReqAuthToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required!"})
		c.Abort()
		return
	}

	// extract token from the header, assuming the format is "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format!"})
		c.Abort()
		return
	}

	token := d.Token{}
	if token.Get(sgl.DATABASE, tokenString) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find TOKEN in database!"})
		c.Abort()
		return
	}

	user := d.User{}
	if user.GetByID(sgl.DATABASE, token.UserID) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find USER in database!"})
		c.Abort()
		return
	}

	c.Set("user", &user)
	c.Next()
}
