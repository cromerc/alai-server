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

func ListObjectState(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectStates []models.ObjectState

	queryParams := request.URL.Query()

	limit, offset, err := utils.GetLimitOffset(queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	filters := []string{
		"name",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.ObjectState{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&objectStates)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		var objectStatesPublic []models.ObjectStatePublic
		for _, objectstate := range objectStates {
			objectStatesPublic = append(objectStatesPublic, models.ObjectStatePublic{ObjectState: objectstate})
		}

		json.NewEncoder(writer).Encode(objectStatesPublic)
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
		json.NewEncoder(writer).Encode(models.ObjectStatePublic{ObjectState: objectState})
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

	result := gdb.Unscoped().Delete(&objectState)
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
