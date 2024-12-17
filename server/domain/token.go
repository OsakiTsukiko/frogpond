package domain

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID         uint       `gorm:"primaryKey"`
	UserID     uint       `gorm:"not null;index"` // foreign key to users table
	Token      string     `gorm:"unique;not null"`
	ClientName string     `gorm:"not null"`
	ExpiresAt  *time.Time // optional expiration (maybe in the future?)
	CreatedAt  time.Time  // automatic timestamp by gorm
}

func (Token) TableName() string {
	return "tokens"
}

// DATABASE METHODS

// TODO: Maybe suffix all database methods with DB?

// add token to database
func (token *Token) Create(db *gorm.DB) error {
	return db.Create(token).Error
}

func (Token) Delete(token string, db *gorm.DB) error {
	return db.Where("token = ?", token).Delete(Token{}).Error
}

func (Token) Exists(db *gorm.DB, token string) (bool, error) {
	var count int64
	err := db.Model(&Token{}).Where("token = ?", token).Count(&count).Error
	return count > 0, err
}

func (token *Token) Get(db *gorm.DB, tokenString string) error {
	return db.Where("token = ?", tokenString).First(token).Error
}
