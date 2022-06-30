package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	ErrorMessage string `json:"error_message"`
}

func JSONErrorOutput(writer http.ResponseWriter, status int, msg string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(ErrorMessage{ErrorMessage: msg})
}
