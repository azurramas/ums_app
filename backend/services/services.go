package services

import (
	"encoding/json"
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
