package main

import (
	"encoding/json"
	"net/http"
)


func respondWithError(w http.ResponseWriter, code int, msg string) {
	payload := map[string]string{
		"error": msg,
	}

	respondWithJSON(w, code, payload)
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	body, err := json.Marshal(payload)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("error incoding the request"))
		return		
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}
