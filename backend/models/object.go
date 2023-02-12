package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	ID            uint64      `json:"ID,omitempty" gorm:"primaryKey"`
	FrameID       uint64      `json:"frame_id,omitempty" gorm:"not null"`
	Frame         Frame       `json:"frame" gorm:"not null"`
	ObjectNameID  uint64      `json:"object_name_id" gorm:"not null"`
	ObjectName    ObjectName  `json:"object_name" gorm:"not null"`
	ObjectStateID uint64      `json:"object_state_id" gorm:"not null"`
	ObjectState   ObjectState `json:"object_state" gorm:"not null"`
	Name          string      `json:"name,omitempty" gorm:"-:all"`
	State         string      `json:"state,omitempty" gorm:"-:all"`
	PositionX     float64     `json:"position_x" gorm:"not null"`
	PositionY     float64     `json:"position_y" gorm:"not null"`
	VelocityX     float64     `json:"velocity_x" gorm:"not null"`
	VelocityY     float64     `json:"velocity_y" gorm:"not null"`
}

type ObjectPublic struct {
	Object
	Frame       bool `json:"frame,omitempty"`
	ObjectName  bool `json:"object_name,omitempty"`
	ObjectState bool `json:"object_state,omitempty"`
	Name        bool `json:"name,omitempty"`
	State       bool `json:"state,omitempty"`
	CreatedAt   bool `json:"CreatedAt,omitempty"`
	UpdatedAt   bool `json:"UpdatedAt,omitempty"`
	DeletedAt   bool `json:"DeletedAt,omitempty"`
}
