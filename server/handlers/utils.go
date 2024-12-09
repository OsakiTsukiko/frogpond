package handlers

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func queryFromArray(parameters []string) string {
	var res = ""

	for i, v := range parameters {
		if i == 0 {
			res += "?" + v
		} else {
			res += "&" + v
		}
	}

	return res
}

func isValidUsername(username string) bool {
	// Check if username matches the allowed pattern and length (3-16)
	match, _ := regexp.MatchString("^[a-zA-Z0-9_.]{3,16}$", username)
	return match
}
