package utils

import (
	"encoding/json"
	"inventaris/model"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, success bool, status int, message string, data interface{}) {
	response := model.Response{
		Success: success,
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendJSONResponsePagination(w http.ResponseWriter, success bool,page, limit, totalItems,totalPages, status int, message string, data interface{}) {
	response := model.Response{
		Success: success,
		Page: page,
		Limit: limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
