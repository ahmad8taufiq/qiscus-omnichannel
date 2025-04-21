package response

import (
	"encoding/json"
	"net/http"
)

type JSON = map[string]interface{}

func JSONResponse(w http.ResponseWriter, status int, payload JSON) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	JSONResponse(w, http.StatusOK, JSON{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusBadRequest, JSON{
		"success": false,
		"message": message,
	})
}

func Unauthorized(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusUnauthorized, JSON{
		"success": false,
		"message": message,
	})
}

func NotFound(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusNotFound, JSON{
		"success": false,
		"message": message,
	})
}
