package controllers

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func PostGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	decoded := base64.NewDecoder(base64.StdEncoding, request.Body)

	uncompressed, err := gzip.NewReader(decoded)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer uncompressed.Close()

	decoder := json.NewDecoder(uncompressed)

	var game models.Game

	err = decoder.Decode(&game)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = game.Validate()
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := postGame(game, gdb)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}
}

func postGame(game models.Game, gdb *gorm.DB) *gorm.DB {
	return gdb.Create(&game)
}

func ListGames(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var games []models.Game

	result := listGames(&games, gdb)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(games)
		return
	}
}

func listGames(games *[]models.Game, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Game{}).Order("ID asc").Joins("Player").Joins("Level").Joins("OS").Joins("GodotVersion").Find(&games)
}

func GetGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var game models.Game

	result := getGame(&game, params.ByName("id"), gdb)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "A game with the ID "+params.ByName("id")+" does not exist!")
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(game)
		return
	}
}

func getGame(games *models.Game, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Game{}).Order("ID asc").Joins("Player").Joins("Level").Joins("OS").Find(&games, id)
}
