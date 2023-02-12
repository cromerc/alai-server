package models

import "gorm.io/gorm"

type ObjectState struct {
	gorm.Model
	ID   uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"unique;not null"`
}

type ObjectStatePublic struct {
	ObjectState
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
