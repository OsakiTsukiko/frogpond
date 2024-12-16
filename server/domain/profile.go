package domain

import "gorm.io/gorm"

// TODO: HOOK THIS INTO USER
type Profile struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"unique;not null;constraint:OnDelete:CASCADE;"`
	// foreign key to users table
	DisplayName string `gorm:"not null"`
	Bio         string `gorm:""`
	AvatarURL   string `gorm:""`
	BannerURL   string `gorm:""`
}

func (Profile) TableName() string {
	return "profiles"
}

func (profile *Profile) Create(db *gorm.DB) error {
	return db.Create(profile).Error
}
