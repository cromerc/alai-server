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

func ListObjectName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectName []models.ObjectName

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

	result := gdb.Model(&models.ObjectName{}).Order("ID asc").Limit(limit).Offset(offset).Find(&objectName)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(objectName)
	}
}

func GetObjectName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectName models.ObjectName

	result := gdb.Model(&models.ObjectName{}).Find(&objectName, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(objectName)
	}
}

func CreateObjectName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectName models.ObjectName

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&objectName)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Create(&objectName)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateObjectName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectName models.ObjectName

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&objectName)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	objectName.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Updates(&objectName)
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

func DeleteObjectName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var objectName models.ObjectName
	var err error
	objectName.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Delete(&objectName)
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
