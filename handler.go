package main

import "net/http"


func handleErr(w http.ResponseWriter, r *http.Request){
	// whenever we want to throw error for any path operation we use this function there 
	respondWithError(w, 400, "Something went wrong")
}