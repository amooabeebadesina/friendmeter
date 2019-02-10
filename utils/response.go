package utils

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse struct
type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// SendSuccessResponse with status 200
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	resData := SuccessResponse{
		Status: "success",
		Data:   data,
	}
	res, _ := json.Marshal(resData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// SendErrorResponse with status 200
func SendErrorResponse(w http.ResponseWriter, msg string) {
	resData := ErrorResponse{
		Status: "error",
		Msg:    msg,
	}
	res, _ := json.Marshal(resData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
