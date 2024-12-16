package handlers

import (
	"net/http"
	"net/url"

	"github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"email,required"`
	Password string `form:"password" binding:"required"`
}

func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func RegisterPOST(c *gin.Context) {
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

	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		// handle validation errors
		parameters = append(parameters, "error="+url.QueryEscape(err.Error()))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	// validate username
	if !isValidUsername(form.Username) {
		parameters = append(parameters, "error="+url.QueryEscape("Username must be between 3-16 characters and contain only [a-zA-Z0-9_.]!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	// validate password length
	if len(form.Password) < 8 || len(form.Password) > 16 {
		parameters = append(parameters, "error="+url.QueryEscape("Password must be between 8 and 16 characters!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	// Hash Password
	hashedPassword, err := HashPassword(form.Password)
	if err != nil {
		parameters = append(parameters, "error="+url.QueryEscape("Failed to hash password!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	// Create the new user in the database
	user := domain.User{
		Username:     form.Username,
		Email:        form.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := user.Create(sgl.DATABASE); err != nil {
		parameters = append(parameters, "error="+url.QueryEscape("Failed to create user in database!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	err = SessionFromUser(c, &user) // create session cookie
	if err != nil {                 // return error if it fails
		parameters = append(parameters, "error="+url.QueryEscape("Failed to create token!"))
		query := queryFromArray(parameters)
		c.Redirect(http.StatusFound, "/auth/register"+query)
		return
	}

	if has_redirect { // redirect to parameter if exists
		c.Redirect(http.StatusFound, redirect)
		return
	}
	c.Redirect(http.StatusFound, "/") // redirect to home
	return
}
