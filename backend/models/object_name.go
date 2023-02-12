package models

import "gorm.io/gorm"

type ObjectName struct {
	gorm.Model
	ID   uint64 `json:"ID,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"unique;not null"`
}

type ObjectNamePublic struct {
	ObjectName
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
