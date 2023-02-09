package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/utils"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

const JWTContextKey contextKey = "JWTClaims"

func Authenticate(handle httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			utils.JSONErrorOutput(writer, http.StatusUnauthorized, errors.New("no token received").Error())
			return
		}
		tokenString := splitToken[1]

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, JWTContextKey, claims)
		request = request.WithContext(ctx)

		handle(writer, request, params)
	}
}
