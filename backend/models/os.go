package models

import "gorm.io/gorm"

type OS struct {
	gorm.Model
	ID   uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"unique;size:8;not null"`
}

type OSPublic struct {
	OS
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
