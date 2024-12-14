package database

import (
	"github.com/OsakiTsukiko/frogpond/server/domain"
	"gorm.io/gorm"
)

func AddToken(userID uint, token string, client_name string, db *gorm.DB) error {
	newToken := domain.Token{
		UserID:     userID,
		Token:      token,
		ClientName: client_name,
	}
	return db.Create(&newToken).Error
}

func GetUserTokens(userID uint, db *gorm.DB) ([]domain.Token, error) {
	var tokens []domain.Token
	if err := db.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func TokenExists(token string, db *gorm.DB) (bool, error) {
	var count int64
	err := db.Model(&domain.Token{}).Where("token = ?", token).Count(&count).Error
	return count > 0, err
}

func RemoveToken(token string, db *gorm.DB) error {
	return db.Where("token = ?", token).Delete(&domain.Token{}).Error
}

func RemoveAllTokensForUser(userID uint, db *gorm.DB) error {
	return db.Where("user_id = ?", userID).Delete(&domain.Token{}).Error
}
