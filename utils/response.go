package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse struct
type JSONResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// SendSuccessResponse with status 200
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	resData := JSONResponse{
		Status: "success",
		Data:   data,
	}
	res, _ := json.Marshal(resData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// SendErrorResponse with status 200
func SendErrorResponse() {

}
