package models

import "gorm.io/gorm"

type ObjectName struct {
	gorm.Model
	ID   uint64 `json:"ID" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
}
