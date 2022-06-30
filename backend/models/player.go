package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID    uint64 `json:"ID" gorm:"primaryKey"`
	RUT   string `json:"rut" gorm:"unique;size:9;not null"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}
