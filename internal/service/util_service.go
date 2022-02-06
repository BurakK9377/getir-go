package service

import (
	"encoding/json"
	"getir/internal/models/record"
	"net/http"
)

// GetError : This is service function to prepare error model.
func getError(err error, w http.ResponseWriter, httpStatusCode int) {
	response := record.ErrorResponse{
		Msg:  err.Error(),
		Code: httpStatusCode,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.Code)
	w.Write(message)
}
