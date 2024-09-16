package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, marshalErr := json.Marshal(payload)

	if marshalErr != nil {
		log.Println("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	errorStruct := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	error, marshalErr := json.Marshal(errorStruct)
	if marshalErr != nil {
		log.Println("Failed to marshal JSON ERROR response: %v", errorStruct)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(error)
}
