package utils

import (
	"encoding/json"
	"net/http"
)
type JsonResponse struct {
	Code    int   `json:"code"`
	Status  string   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, payload JsonResponse, headers ...http.Header) error {
	out, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if len(headers)>0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Code)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error,  status string, code ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = code[0]
	}

	var payload JsonResponse
	payload.Status = status
	payload.Code = statusCode

	if err != nil {
		payload.Message = err.Error()  
	} else {
		payload.Message ="Something went wrong"
	}

	return WriteJSON(w, statusCode, payload)
}