package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mhmd-hariri/go-bookstore/pkg/config"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

func init() {

	config.DB.AutoMigrate(&User{})
}

func GetUserByUsername(username string) (*User, *gorm.DB, error) {
	var user User
	db := config.DB.Where("USERNAME=?", username).Find(&user)
	if db.Error != nil {
		return nil, db, db.Error
	}
	return &user, db, nil
}

// CreateUser creates a new user with a hashed password
func (user *User) CreateUser() (*User, error) {

	config.DB.NewRecord(user)
	db := config.DB.Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}
