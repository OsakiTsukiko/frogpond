package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

// DATABASE METHODS

// TODO: Maybe suffix all database methods with DB?

// create new user
func (user *User) Create(db *gorm.DB) error {
	return db.Create(user).Error
}

// delete user
func (user *User) Delete(db *gorm.DB) error {
	return db.Delete(user).Error
}

// find by user id
func (User) GetByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	return &user, err
}

// find by username
func (User) GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// no find by email required for now.

// update user (based on id)
func (user *User) Update(db *gorm.DB) error {
	return db.Save(user).Error
}

// list all users in database
// TODO: ADD PAGINATION
func (User) ListAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

// checks if username is taken
func (User) IsUsernameTaken(db *gorm.DB, username string) (bool, error) {
	var count int64
	err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// checks if email is taken
func (User) IsEmailTaken(db *gorm.DB, email string) (bool, error) {
	var count int64
	err := db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// validate username password (hash) combination
func (User) AuthenticateUser(db *gorm.DB, username, password string) (*User, error) {
	var user User

	// get user from database
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// compare the provided password with the stored password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
