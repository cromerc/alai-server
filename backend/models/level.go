package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	ID   uint64 `json:"ID" gorm:"primaryKey"`
	Name string `json:"rut" gorm:"unique;not null"`
}
