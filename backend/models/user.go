package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	Name        string `json:"name,omitempty" gorm:"not null"`
	Username    string `json:"username,omitempty" gorm:"unique; not null"`
	Email       string `json:"email,omitempty" gorm:"unique;not null"`
	Password    string `json:"password,omitempty" gorm:"not null"`
	NewPassword string `json:"new_password,omitempty" gorm:"-:all"`
}

type UserPublic struct {
	User
	Password    bool `json:"password,omitempty"`
	NewPassword bool `json:"new_password,omitempty"`
	CreatedAt   bool `json:"CreatedAt,omitempty"`
	UpdatedAt   bool `json:"UpdatedAt,omitempty"`
	DeletedAt   bool `json:"DeletedAt,omitempty"`
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
