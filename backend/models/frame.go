package models

import (
	"gorm.io/gorm"
)

type Frame struct {
	gorm.Model
	ID          uint64   `json:"ID" gorm:"primaryKey"`
	GameID      uint64   `json:"game_id" gorm:"not null"`
	Game        Game     `json:"game" gorm:"not null"`
	Coins       uint64   `json:"coins" gorm:";not null"`
	Points      uint64   `json:"points" gorm:"not null"`
	FPS         uint8    `json:"fps" gorm:"not null"`
	ElapsedTime uint64   `json:"elapsed_time" gorm:"not null"`
	Objects     []Object `json:"objects"`
}
