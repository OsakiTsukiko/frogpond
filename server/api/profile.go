package api

import (
	"net/http"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	user_any, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to access USER context!",
		})
		return
	}
	user := user_any.(*d.User)

	profile := &d.Profile{}
	if profile.ForUser(sgl.DATABASE, user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to access PROFILE from database!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      profile.UserID,
		"display_name": profile.DisplayName,
		"bio":          profile.Bio,
		"avatar_url":   profile.AvatarURL,
		"banner_url":   profile.BannerURL,
	})
}
