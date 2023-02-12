package controllers

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func ListGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var games []models.Game

	queryParams := request.URL.Query()

	limit, offset, err := utils.GetLimitOffset(queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	filters := []string{
		"player_id",
		"level_id",
		"os_id",
		"godot_version_id",
		"processor_count",
		"screen_count",
		"screen_dpi",
		"screen_size",
		"machine_id",
		"locale",
		"game_version",
		"won",
		"timestamp",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.Game{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&games)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		var gamesPublic []models.GamePublic
		for _, game := range games {
			gamesPublic = append(gamesPublic, models.GamePublic{Game: game})
		}

		json.NewEncoder(writer).Encode(gamesPublic)
		return
	}
}

func GetGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var game models.Game

	result := gdb.Model(&models.Game{}).Order("ID asc").Find(&game, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(models.GamePublic{Game: game})
		return
	}
}

func CreateGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	result := gdb.Create(&game)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}
}

func UpdateGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var game models.Game

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&game)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	game.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&game)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func DeleteGame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var game models.Game
	var err error
	game.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Unscoped().Delete(&game)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}
