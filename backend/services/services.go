package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteError(w http.ResponseWriter, err error) {
	var message map[string]interface{}
	var statusCode int

	switch err {
	case sql.ErrNoRows:
		message = map[string]interface{}{
			"error": "Record does not exist",
			"desc":  err.Error(),
		}

		statusCode = http.StatusNotFound

	default:
		message = map[string]interface{}{
			"error": "Error while processing request",
			"desc":  err.Error(),
		}

		statusCode = http.StatusInternalServerError

	}

	fmt.Println("Error:", err.Error())
	WriteResponse(w, message, statusCode)
}
