package models

import (
	"gorm.io/gorm"
)

type Frame struct {
	gorm.Model
	ID          uint64   `json:"ID,omitempty" gorm:"primaryKey"`
	GameID      uint64   `json:"game_id,omitempty" gorm:"not null"`
	Game        Game     `json:"game" gorm:"not null"`
	Coins       uint64   `json:"coins" gorm:";not null"`
	Points      uint64   `json:"points,omitempty" gorm:"not null"`
	FPS         uint8    `json:"fps,omitempty" gorm:"not null"`
	ElapsedTime uint64   `json:"elapsed_time" gorm:"not null"`
	Objects     []Object `json:"objects,omitempty"`
}

type FramePublic struct {
	Frame
	Game      bool `json:"game,omitempty"`
	CreatedAt bool `json:"CreatedAt,omitempty"`
	UpdatedAt bool `json:"UpdatedAt,omitempty"`
	DeletedAt bool `json:"DeletedAt,omitempty"`
}
