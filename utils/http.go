package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)

	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}

func ErrorResponse(writer http.ResponseWriter, error string, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := struct {
		Message string `json:"message"`
	}{error}
	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}
