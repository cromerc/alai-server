package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func ListUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var users []models.User

	queryParams := request.URL.Query()

	limit, offset, err := utils.GetLimitOffset(queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	filters := []string{
		"name",
		"username",
		"email",
		"password",
	}

	whereClause, err := utils.GenerateWhereFilter(filters, queryParams)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(&models.User{}).Where(whereClause).Order("ID asc").Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		var usersPublic []models.UserPublic
		for _, user := range users {
			usersPublic = append(usersPublic, models.UserPublic{User: user})
		}

		json.NewEncoder(writer).Encode(usersPublic)
	}
}

func GetUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var user models.User

	result := gdb.Model(&models.User{}).Find(&user, params.ByName("id"))
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(models.UserPublic{User: user})
	}
}

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var user models.User

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&user)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	user.HashPassword(user.Password)

	result := gdb.Create(&user)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	claims := request.Context().Value(middlewares.JWTContextKey).(*utils.JWTClaim)
	username := claims.Username

	var user models.User

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&user)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	user.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	if user.NewPassword != "" {
		var tmpUser models.User

		result := gdb.Where(&models.User{Username: username}).Find(&tmpUser)
		if result.Error != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
			return
		} else if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		// If the logged in user and the modified user are no the same, password can't be changed
		if tmpUser.ID != user.ID {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("only the same user may change password").Error())
			return
		}

		err = tmpUser.CheckPassword(user.Password)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("incorrect user or password").Error())
			return
		}

		user.HashPassword(user.NewPassword)
	} else {
		user.Password = ""
	}

	result := gdb.Updates(&user)
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

func DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var user models.User
	user.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := gdb.Unscoped().Delete(&user)
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
