package database

import (
	"fmt"

	"github.com/OsakiTsukiko/frogpond/server/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// returns user database or error
// keep all errors custom as upstream
// returns them to the user
func GetUserFromDatabase(username, password string, db *gorm.DB) (*domain.User, error) {
	var user domain.User
	// query the user by username
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("Invalid Credentials")
		// keep this as Invalid Credentials for security
	}

	// compare the provided password with the stored password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	return &user, nil
}
