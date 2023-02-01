package middlewares

import (
	"net/http"
)

func Cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		origin := request.Header.Get("Origin")

		switch origin {
		case "http://localhost:5173", "http://localhost", "https://alai.cromer.cl":
			writer.Header().Set("Access-Control-Allow-Origin", origin)
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		handler.ServeHTTP(writer, request)
	})
}
