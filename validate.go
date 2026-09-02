package main

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
)
func validateChrip(w http.ResponseWriter, r *http.Request) {
    type returnVals struct {
        Body string `json:"body"`
    }

    decoder := json.NewDecoder(r.Body)
    jreq := returnVals{}

    err := decoder.Decode(&jreq)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request")
        return
    }
		len:= utf8.RuneCountInString(jreq.Body)
		if len>140{
			respondWithError(w, http.StatusBadRequest, "Chirp is too long")	
			return
		}
		cleaned:= replaceBadWords(jreq.Body)
		respondWithJSON(w, http.StatusOK, map[string]string{"cleaned_body": cleaned,})
}

