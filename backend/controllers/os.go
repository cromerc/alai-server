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

func ListOS(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var os []models.OS

	result := gdb.Model(&models.OS{}).Order("ID asc").Find(&os)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(os)
	}
}

func GetOS(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var os models.OS

	result := gdb.Model(&models.OS{}).Find(&os, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(os)
	}
}

func CreateOS(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var os models.OS

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&os)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Create(&os)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateOS(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var os models.OS

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&os)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	os.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&os)
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

func DeleteOS(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var os models.OS
	var err error
	os.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Delete(&os)
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
