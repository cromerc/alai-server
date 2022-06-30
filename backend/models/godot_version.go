package models

import "gorm.io/gorm"

type GodotVersion struct {
	gorm.Model
	ID     uint64 `json:"ID" gorm:"primaryKey"`
	Major  uint8  `json:"major" gorm:"not null"`
	Minor  uint8  `json:"minor" gorm:"not null"`
	Patch  uint8  `json:"patch" gorm:"not null"`
	Hex    uint64 `json:"hex" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
	Build  string `json:"build" gorm:"not null"`
	Year   uint16 `json:"year" gorm:"not null"`
	Hash   string `json:"hash" gorm:"unique;size:40;not null"`
	String string `json:"string" gorm:"unique;not null"`
}
