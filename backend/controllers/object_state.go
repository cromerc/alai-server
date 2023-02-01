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

func ListObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectState []models.ObjectState

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

	filters := []string{
		"name",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.ObjectState{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&objectState)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(objectState)
	}
}

func GetObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectState models.ObjectState

	result := gdb.Model(&models.ObjectState{}).Find(&objectState, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(objectState)
	}
}

func CreateObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectState models.ObjectState

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&objectState)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Create(&objectState)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectState models.ObjectState

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&objectState)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	objectState.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&objectState)
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

func DeleteObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectState models.ObjectState
	var err error
	objectState.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Delete(&objectState)
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
