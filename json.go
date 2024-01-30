package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError function handles HTTP error responses by sending a JSON-encoded error message
// It takes a ResponseWriter, HTTP status code, and an error message as input parameters
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// If the HTTP status code is greater than 499 (indicating a server error),
	// log the error message to the console
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	// Define a struct for the error response with a single field 'Error'
	type errResponse struct {
		Error string `json:"error"`
	}

	// Create a JSON-encoded error response using the respondWithJSON function
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// respondWithJSON function sends a JSON response with the provided payload and HTTP status code
// It takes a ResponseWriter, HTTP status code, and a payload (interface{}) as input parameters
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON format
	dat, err := json.Marshal(payload)
	if err != nil {
		// If there is an error during JSON marshaling, log the error and return a 500 Internal Server Error
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	// Set the Content-Type header to indicate that the response contains JSON data
	w.Header().Add("Content-Type", "application/json")

	// Set the HTTP status code for the response
	w.WriteHeader(code)

	// Write the JSON-encoded data to the ResponseWriter
	w.Write(dat)
}
