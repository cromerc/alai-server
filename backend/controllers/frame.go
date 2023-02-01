package controllers

import (
	"encoding/json"
	"math"
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

	var frame []models.Frame

	queryParams := request.URL.Query()

	limit := 50
	if queryParams.Get("limit") != "" {
		var err error
		limit, err = strconv.Atoi(queryParams.Get("limit"))
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
			return
		}
		limit = int(math.Min(float64(500), float64(limit)))
		limit = int(math.Max(float64(1), float64(limit)))
	}

	offset := 0
	if queryParams.Get("offset") != "" {
		var err error
		offset, err = strconv.Atoi(queryParams.Get("offset"))
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
			return
		}
		offset = int(math.Min(float64(9223372036854775807), float64(offset)))
		offset = int(math.Max(float64(0), float64(offset)))
	}

	result := gdb.Model(&models.Frame{}).Order("ID asc").Limit(limit).Offset(offset).Find(&frame)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(frame)
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
		json.NewEncoder(writer).Encode(frame)
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

	result := gdb.Delete(&frame)
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
