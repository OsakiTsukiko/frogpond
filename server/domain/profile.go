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

	User User `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

func (Profile) TableName() string {
	return "profiles"
}

func (profile *Profile) ForUser(db *gorm.DB, user *User) error {
	return db.Where("user_id = ?", user.ID).First(profile).Error
}
