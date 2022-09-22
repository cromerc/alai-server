package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID             uint64       `json:"ID" gorm:"primaryKey"`
	PlayerID       uint64       `json:"player_id"`
	Player         Player       `json:"player"`
	LevelID        uint64       `json:"level_id" gorm:"not null"`
	Level          Level        `json:"level" gorm:"not null"`
	OSID           uint64       `json:"os_id" gorm:"not null"`
	OS             OS           `json:"os" gorm:"not null"`
	GodotVersionID uint64       `json:"godot_version_id" gorm:"not null"`
	GodotVersion   GodotVersion `json:"godot_version" gorm:"not null"`
	ProcessorCount uint64       `json:"processor_count" gorm:"not null"`
	ScreenCount    uint8        `json:"screen_count" gorm:"not null"`
	ScreenDPI      uint8        `json:"screen_dpi" gorm:"not null"`
	ScreenSize     string       `json:"screen_size" gorm:"not null"`
	MachineId      string       `json:"machine_id" gorm:"not null"`
	Locale         string       `json:"locale" gorm:"not null"`
	GameVersion    string       `json:"game_version" gorm:"not null"`
	Won            bool         `json:"won" gorm:"not null"`
	Timestamp      uint64       `json:"timestamp" gorm:"not null"`
	Frames         []Frame      `json:"frames"`
}

func (game *Game) Validate() error {
	if len(strings.TrimSpace(game.MachineId)) == 0 {
		return errors.New("empty machine id")
	}
	return nil
}

// Cache the results of the queries here
// The object states and names should not be deleted so this should always be valid while running
var cachedStateNames = make(map[string]uint64)
var cachedObjectNames = make(map[string]uint64)

func (game *Game) BeforeCreate(tx *gorm.DB) error {
	// Use the same player ID if the RUT is already in the DB
	tx.Model(Player{}).Where(&Player{RUT: game.Player.RUT}).Find(&game.Player)

	tx.Model(GodotVersion{}).Where(&GodotVersion{String: game.GodotVersion.String}).Find(&game.GodotVersion)

	for frameIndex, frame := range game.Frames {
		for objectIndex := range frame.Objects {
			game.Frames[frameIndex].Objects[objectIndex].ObjectState.Name = game.Frames[frameIndex].Objects[objectIndex].State
			game.Frames[frameIndex].Objects[objectIndex].ObjectName.Name = game.Frames[frameIndex].Objects[objectIndex].Name

			// Use the existing state names in the database if they exist
			if ID, ok := cachedStateNames[game.Frames[frameIndex].Objects[objectIndex].ObjectState.Name]; ok {
				// The name is cached, no need to query the database
				game.Frames[frameIndex].Objects[objectIndex].ObjectStateID = ID
				game.Frames[frameIndex].Objects[objectIndex].ObjectState.ID = ID
			} else {
				var state ObjectState
				result := tx.Model(ObjectState{}).Where(&ObjectState{Name: game.Frames[frameIndex].Objects[objectIndex].ObjectState.Name}).Find(&state)
				if result.RowsAffected == 0 {
					// Not in the database, so let's create it
					tx.Create(&game.Frames[frameIndex].Objects[objectIndex].ObjectState)
					game.Frames[frameIndex].Objects[objectIndex].ObjectStateID = game.Frames[frameIndex].Objects[objectIndex].ObjectState.ID
					cachedStateNames[game.Frames[frameIndex].Objects[objectIndex].ObjectState.Name] = game.Frames[frameIndex].Objects[objectIndex].ObjectState.ID
				} else {
					// It is in the database, so use that
					game.Frames[frameIndex].Objects[objectIndex].ObjectStateID = state.ID
					game.Frames[frameIndex].Objects[objectIndex].ObjectState.ID = state.ID
					cachedStateNames[game.Frames[frameIndex].Objects[objectIndex].ObjectState.Name] = state.ID
				}
			}

			// Use the existing object names in the database if they exist
			if ID, ok := cachedObjectNames[game.Frames[frameIndex].Objects[objectIndex].ObjectName.Name]; ok {
				// The name is cached, no need to query the database
				game.Frames[frameIndex].Objects[objectIndex].ObjectNameID = ID
				game.Frames[frameIndex].Objects[objectIndex].ObjectName.ID = ID
			} else {
				var objectName ObjectName
				result := tx.Model(ObjectName{}).Where(&ObjectName{Name: game.Frames[frameIndex].Objects[objectIndex].ObjectName.Name}).Find(&objectName)
				if result.RowsAffected == 0 {
					// Not in the database, so let's create it
					tx.Create(&game.Frames[frameIndex].Objects[objectIndex].ObjectName)
					game.Frames[frameIndex].Objects[objectIndex].ObjectNameID = game.Frames[frameIndex].Objects[objectIndex].ObjectName.ID
					cachedObjectNames[game.Frames[frameIndex].Objects[objectIndex].ObjectName.Name] = game.Frames[frameIndex].Objects[objectIndex].ObjectName.ID
				} else {
					// It is in the database, so use that
					game.Frames[frameIndex].Objects[objectIndex].ObjectNameID = objectName.ID
					game.Frames[frameIndex].Objects[objectIndex].ObjectName.ID = objectName.ID
					cachedObjectNames[game.Frames[frameIndex].Objects[objectIndex].ObjectName.Name] = objectName.ID
				}
			}
		}
	}
	return nil
}
