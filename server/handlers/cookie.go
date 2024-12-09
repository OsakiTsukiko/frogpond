package handlers

import (
	"fmt"
	"time"

	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// converts session cookie (containing jwt token) to user
// (username, email, ok)
// TODO: is email even useful? usernames are unmutable..
func UserFromSession(c *gin.Context) (string, string, bool) {
	// get the session cookie from the request
	cookie, err := c.Cookie("session")
	if err != nil { // cookie doesn't exist
		return "", "", false
	}

	// parse the JWT token from the cookie
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// validate the token signing method (HS256 => HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// the return value is never actually used..
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return the secret key used for signing the token
		return []byte(sgl.CFG.Server.JWTSecretKey), nil
	})

	// check if token is valid
	if err != nil || !token.Valid {
		return "", "", false
	}

	// extract user information from the JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", false
	}

	// extract username and email from claims (or return empty if not found)
	username, usernameOk := claims["username"].(string)
	email, emailOk := claims["email"].(string)

	if !usernameOk || !emailOk {
		return "", "", false
	}

	// eeturn the username and email if everything is valid
	return username, email, true
}

func SessionFromUser(c *gin.Context, username string, email string) error {
	// create JWT token (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24 * 8).Unix(),
		// 8 day expiration
	})

	tokenString, err := token.SignedString([]byte(sgl.CFG.Server.JWTSecretKey))
	if err != nil {
		return err
	}

	// Set session cookie
	c.SetCookie(
		"session",
		tokenString,
		3600*24*8,               // 8 days
		"/",                     // make global
		sgl.CFG.Server.Domain,   // domain
		sgl.CFG.Server.UseHTTPS, // secure, https
		true,                    // HttpOnly (no js?)
	)

	return nil
}

func ClearSession(c *gin.Context) {
	c.SetCookie(
		"session",
		"",
		-1, // expire the cookie immediately
		"/",
		sgl.CFG.Server.Domain,
		sgl.CFG.Server.UseHTTPS,
		true,
	)
}
