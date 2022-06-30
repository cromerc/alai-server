package models

import "gorm.io/gorm"

type OS struct {
	gorm.Model
	ID   uint64 `json:"ID" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;size:8;not null"`
}
