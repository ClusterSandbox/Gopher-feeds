package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 { 
		log.Printf("Responding with %v error: %v", code, msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	// respondWithJSON()
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// we marshall the payload to json as this is a json REST API

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshall JSON Response: %v", payload)
		log.Printf("Error occurred: %v", err)
		w.WriteHeader(500) // give a status code of 500 
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) // status_code=code that was passed that that json conversion worked fine 
	w.Write(data)


}