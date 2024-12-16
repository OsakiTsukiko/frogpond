package handlers

import (
	"net/http"
	"net/url"

	"github.com/OsakiTsukiko/frogpond/server/domain"
	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginPOST(c *gin.Context) {
	// get redirect parameter from query
	redirect_escaped := c.Query("redirect")
	has_redirect := redirect_escaped != ""
	redirect, err := url.QueryUnescape(redirect_escaped)
	if err != nil {
		has_redirect = false
	}

	var parameters = []string{}
	if has_redirect {
		parameters = append(parameters, "redirect="+redirect_escaped)
	}

	// local database pointer for ease of use
	db := sgl.DATABASE

	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		// handle validation errors
		parameters = append(parameters, "error="+url.QueryEscape(err.Error()))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/login"+query)
		return
	}

	// Get user from database (validate)
	user, err := domain.User.AuthenticateUser(
		d.User{},
		db,
		form.Username,
		form.Password,
	)

	if err != nil {
		// handle validation errors
		parameters = append(parameters, "error="+url.QueryEscape("Invalid Credentials!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/login"+query)
		return
	}

	// create session
	SessionFromUser(c, user.Username, user.Email)

	// redirect after login
	if has_redirect { // redirect to parameter if exists
		c.Redirect(http.StatusFound, redirect)
		return
	}
	// redirect to homepage after login
	c.Redirect(http.StatusFound, "/")
	return
}
