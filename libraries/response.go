package libraries

import (
	"encoding/json"
	"net/http"
)

var result result_json

type result_json struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, m string, code int) {
	result = result_json{}
	result.Code = code
	result.Message = m
	printResponse(w)
}

func SetData(w http.ResponseWriter, data interface{}, code int) {
	result = result_json{}
	result.Code = code
	result.Data = data
	printResponse(w)
}

func printResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Code)
	json.NewEncoder(w).Encode(result)
}
