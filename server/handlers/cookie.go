package handlers

import (
	"fmt"
	"time"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	sgl "github.com/OsakiTsukiko/frogpond/server/singleton"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// SESSION MANAGEMENT

// converts session cookie (containing jwt token) to userid
func UserFromSession(c *gin.Context, db *gorm.DB) (*d.User, bool) {
	cookie, err := c.Cookie("session")
	if err != nil {
		return nil, false
	}

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
		return nil, false
	}

	// extract user id from the JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	userid, useridOK := claims["userid"].(float64)

	if !useridOK {
		return nil, false
	}

	user := &d.User{}
	if err := user.GetByID(db, uint(userid)); err != nil {
		return nil, false
	}
	return user, true
}

// converts userid to session cookie
func SessionFromUser(c *gin.Context, user *d.User) error {
	// create JWT token (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": float64(user.ID),
		"exp":    time.Now().Add(time.Hour * 24 * 8).Unix(),
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

// remove any session cookie
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
