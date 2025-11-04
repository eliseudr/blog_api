package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Success returns a success response with status code and data
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SuccessResponse{
		Code:    statusCode,
		Success: true,
		Data:    data,
	})
}

// Error returns an error response with status code and error message
func Error(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    statusCode,
		Success: false,
		Error:   message,
	})
}
