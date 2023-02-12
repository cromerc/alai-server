package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func ListFrame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var frames []models.Frame

	queryParams := request.URL.Query()

	limit, offset, err := utils.GetLimitOffset(queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	filters := []string{
		"game_id",
		"coins",
		"points",
		"fps",
		"elapsed_time",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.Frame{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&frames)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		var framesPublic []models.FramePublic
		for _, frame := range frames {
			framesPublic = append(framesPublic, models.FramePublic{Frame: frame})
		}

		json.NewEncoder(writer).Encode(framesPublic)
	}
}

func GetFrame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var frame models.Frame

	result := gdb.Model(&models.Frame{}).Find(&frame, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(models.FramePublic{Frame: frame})
	}
}

func CreateFrame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var frame models.Frame

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&frame)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Create(&frame)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateFrame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var frame models.Frame

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&frame)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	frame.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&frame)
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

func DeleteFrame(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var frame models.Frame
	var err error
	frame.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Unscoped().Delete(&frame)
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
