package domain

import "gorm.io/gorm"

// TODO: HOOK THIS INTO USER
type Profile struct {
	UserID      uint   `gorm:"unique;not null"` // foreign key to users table
	DisplayName string `gorm:"not null"`
	Description string `gorm:""`
	AvatarURL   string `gorm:""`
	BannerURL   string `gorm:""`
}

func (Profile) TableName() string {
	return "profiles"
}

func (profile *Profile) Create(db *gorm.DB) error {
	return db.Create(profile).Error
}
