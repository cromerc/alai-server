package models

import "gorm.io/gorm"

type GodotVersion struct {
	gorm.Model
	ID     uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	Major  uint8  `json:"major,omitempty" gorm:"not null"`
	Minor  uint8  `json:"minor" gorm:"not null"`
	Patch  uint8  `json:"patch" gorm:"not null"`
	Hex    uint64 `json:"hex,omitempty" gorm:"not null"`
	Status string `json:"status,omitempty" gorm:"not null"`
	Build  string `json:"build,omitempty" gorm:"not null"`
	Year   uint16 `json:"year,omitempty" gorm:"not null"`
	Hash   string `json:"hash,omitempty" gorm:"unique;size:40;not null"`
	String string `json:"string,omitempty" gorm:"unique;not null"`
}

type GodotVersionPublic struct {
	GodotVersion
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
