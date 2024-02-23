package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(resWriter http.ResponseWriter, message string, responseCode int) {
	response := Error{
		Code:    responseCode,
		Message: message,
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(responseCode)

	json.NewEncoder(resWriter).Encode(response)
}

var RequestErrorHandler = func(w http.ResponseWriter, error error) {
	writeError(w, error.Error(), http.StatusBadRequest)
}

var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
}
