package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/database"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/models"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"
	"github.com/julienschmidt/httprouter"
)

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var receivedUser models.User
	var user models.User

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&receivedUser)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := gdb.Model(models.User{}).Where(&models.User{Username: receivedUser.Username}).Find(&user)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("incorrect user or password").Error())
		return
	}

	err = user.CheckPassword(receivedUser.Password)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("incorrect user or password").Error())
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
