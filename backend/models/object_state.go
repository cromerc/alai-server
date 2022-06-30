package models

import "gorm.io/gorm"

type ObjectState struct {
	gorm.Model
	ID   uint64 `json:"ID" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
}
