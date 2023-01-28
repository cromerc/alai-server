package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var user models.User

	gdb.Model(models.User{}).Where(&models.User{Username: request.FormValue("username")}).Find(&user)

	err := user.CheckPassword(request.FormValue("password"))
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("incorrect password").Error())
		return
	}

	type Token struct {
		Token string `json:"token"`
	}

	tokenString, err := utils.GenerateJWT(user.Email, user.Username)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	token := Token{Token: tokenString}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(token)
}

func AuthenticateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}

func ListUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var users []models.User

	result := gdb.Model(&models.User{}).Order("ID asc").Find(&users)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		for i := range users {
			users[i].Password = ""
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(users)
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
		user.Password = ""
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(user)
	}
}

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	user := models.User{
		Name:  request.FormValue("name"),
		Email: request.FormValue("email"),
	}
	user.HashPassword(request.FormValue("password"))

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

	var user models.User

	userID, _ := strconv.ParseUint(params.ByName("id"), 10, 64)

	gdb.Model(models.User{}).Where(&models.User{ID: userID}).Find(&user)

	if request.FormValue("password") != "" {
		var oldPassword = request.FormValue("old_password")
		err := user.CheckPassword(oldPassword)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("incorrect password").Error())
			return

		} else {
			user.HashPassword(request.FormValue("password"))
		}
	}

	user.Name = request.FormValue("name")
	user.Email = request.FormValue("email")

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

	result := gdb.Delete(&user)
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
