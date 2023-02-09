package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint64 `json:"ID" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Username    string `json:"username" gorm:"unique; not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"password,omitempty" gorm:"not null"`
	NewPassword string `json:"new_password,omitempty" gorm:"-:all"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
