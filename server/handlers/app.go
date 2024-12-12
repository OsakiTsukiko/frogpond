package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Parameters
client_name: string, required
website: string, optional
*/

type AppForm struct {
	ClientName string `form:"client_name" binding:"required"`
	Website    string `form:"website"` // not required
}

func AppGET(c *gin.Context) {
	client_name := c.Query("client_name")
	has_client_name := client_name != ""
	website := c.Query("website")
	has_website := website != ""

	fmt.Print("AAAA", client_name, has_client_name, website, has_website, "AAAA")

	if !has_client_name {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Invalid URL! For help, Contact the Client Mentainer.",
		})
		return
	}

	var w, hw string
	if has_website {
		w = website
		hw = "nop"
	} else {
		w = ""
		hw = "display-none"
	}

	username, email, ok := UserFromSession(c)
	if !ok {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Corrupted user data!",
			// this should be UNREACHABLE thanks to
			// the ReqAuthMiddleware
		})
		return
	}

	c.HTML(http.StatusOK, "app.html", gin.H{
		"client_name": client_name,
		"website":     w,
		"has_website": hw,
		"username":    username,
		"email":       email,
	})
}

func AppPOST(c *gin.Context) {
	var form AppForm
	if err := c.ShouldBind(&form); err != nil {
		// handle validation errors
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Invalid POST Request!",
		})
		return
	}

	token, err := generateToken()
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Error generating token!",
		})
		return
	}

	c.HTML(http.StatusOK, "error.html", gin.H{
		"error": "Token: " + token,
	})
	return
}

func generateToken() (string, error) {
	bytes := make([]byte, 32) // 256-bit token
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
