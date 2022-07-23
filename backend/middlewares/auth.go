package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func Authenticate(handle httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			utils.JSONErrorOutput(writer, http.StatusUnauthorized, errors.New("no token received").Error())
			return
		}
		tokenString := splitToken[1]

		err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusUnauthorized, err.Error())
			return
		}

		handle(writer, request, params)
	}
}
