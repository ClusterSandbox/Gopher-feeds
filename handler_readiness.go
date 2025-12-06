package main

import (
	"net/http"
)

// the http.Request curresponds to the incoming http request
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
	
}