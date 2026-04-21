package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func OK(w http.ResponseWriter, data interface{}) {
	writeJSON(w, http.StatusOK, Response{Success: true, Data: data})
}

func Created(w http.ResponseWriter, data interface{}) {
	writeJSON(w, http.StatusCreated, Response{Success: true, Data: data})
}

func Message(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusOK, Response{Success: true, Message: message})
}

func BadRequest(w http.ResponseWriter, err string) {
	writeJSON(w, http.StatusBadRequest, ErrorResponse{Success: false, Error: err})
}

func Unauthorized(w http.ResponseWriter, err string) {
	writeJSON(w, http.StatusUnauthorized, ErrorResponse{Success: false, Error: err})
}

func Forbidden(w http.ResponseWriter) {
	writeJSON(w, http.StatusForbidden, ErrorResponse{Success: false, Error: "insufficient permissions"})
}

func NotFound(w http.ResponseWriter, err string) {
	writeJSON(w, http.StatusNotFound, ErrorResponse{Success: false, Error: err})
}

func InternalError(w http.ResponseWriter) {
	writeJSON(w, http.StatusInternalServerError, ErrorResponse{Success: false, Error: "internal server error"})
}

func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
