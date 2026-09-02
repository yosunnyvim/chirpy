package main

import (
	"MODULE_PATH/internal/database"
	"encoding/json"
	"net/http"
	"unicode/utf8"

	"github.com/google/uuid"
)
type createChirpRequest struct {
    Body   string    `json:"body"`
    UserID uuid.UUID `json:"user_id"`
}
func (cfg *apiConfig) createChirp(w http.ResponseWriter, r *http.Request) {


	decoder := json.NewDecoder(r.Body)
	jreq := createChirpRequest{}

	err := decoder.Decode(&jreq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	len := utf8.RuneCountInString(jreq.Body)
	if len > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	cleaned := replaceBadWords(jreq.Body)

	chirp, err := cfg.dbQueries.CreateChirp(r.Context(),database.CreateChirpParams{
		Body: cleaned,
		UserID: jreq.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp")
		return
	}
	responseChirp :=Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body: chirp.Body,
		UserID: chirp.UserID,
	}	
	respondWithJSON(w, http.StatusCreated, responseChirp)
}
