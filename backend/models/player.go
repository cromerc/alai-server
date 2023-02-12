package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID    uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	RUT   string `json:"rut,omitempty" gorm:"unique;size:9;not null"`
	Name  string `json:"name,omitempty" gorm:"not null"`
	Email string `json:"email,omitempty" gorm:"unique;not null"`
}

type PlayerPublic struct {
	Player
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
