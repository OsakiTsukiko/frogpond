package controller

import (
	"errors"

	d "github.com/OsakiTsukiko/frogpond/server/domain"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *d.User) error {
	if user.Create(db) != nil {
		return errors.New("Failed to create USER!")
	}

	profile := d.Profile{
		UserID:      user.ID,
		DisplayName: user.Username,
		Description: "🐸",
	}

	if profile.Create(db) != nil { // should i remove the user?
		_ = user.Delete(db) // PRAY THIS IS NIL
		return errors.New("Failed to create PROFILE!")
	}

	return nil
}
