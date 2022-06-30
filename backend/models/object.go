package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	ID            uint64      `json:"ID" gorm:"primaryKey"`
	FrameID       uint64      `json:"frame_id" gorm:"not null"`
	Frame         Frame       `json:"frame" gorm:"not null"`
	ObjectNameID  uint64      `json:"-" gorm:"not null"`
	ObjectName    ObjectName  `json:"-" gorm:"not null"`
	ObjectStateID uint64      `json:"-" gorm:"not null"`
	ObjectState   ObjectState `json:"-" gorm:"not null"`
	Name          string      `json:"name" gorm:"-:all"`
	State         string      `json:"state" gorm:"-:all"`
	PositionX     float64     `json:"position_x" gorm:"not null"`
	PositionY     float64     `json:"position_y" gorm:"not null"`
	VelocityX     float64     `json:"velocity_x" gorm:"not null"`
	VelocityY     float64     `json:"velocity_y" gorm:"not null"`
}
