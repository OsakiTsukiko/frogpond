package domain

import "time"

type Token struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null;index"` // foreign key to users table
	Token     string     `gorm:"unique;not null"`
	ExpiresAt *time.Time // optional expiration (maybe in the future?)
	CreatedAt time.Time  // automatic timestamp by gorm
}

func (Token) TableName() string {
	return "tokens"
}
