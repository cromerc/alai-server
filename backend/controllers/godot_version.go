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

func ListGodotVersion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var godotVersion []models.GodotVersion

	queryParams := request.URL.Query()

	limit, offset, err := utils.GetLimitOffset(queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	filters := []string{
		"major",
		"minor",
		"patch",
		"hex",
		"status",
		"build",
		"year",
		"hash",
		"string",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.GodotVersion{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&godotVersion)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(godotVersion)
	}
}

func GetGodotVersion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var godotVersion models.GodotVersion

	result := gdb.Model(&models.GodotVersion{}).Find(&godotVersion, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(godotVersion)
	}
}

func CreateGodotVersion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var godotVersion models.GodotVersion

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&godotVersion)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Create(&godotVersion)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateGodotVersion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var godotVersion models.GodotVersion

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&godotVersion)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	godotVersion.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&godotVersion)
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

func DeleteGodotVersion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var godotVersion models.GodotVersion
	var err error
	godotVersion.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Unscoped().Delete(&godotVersion)
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
