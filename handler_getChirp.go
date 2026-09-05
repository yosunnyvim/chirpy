package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirp(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("chirpID")
	uuidId, err := uuid.Parse(strID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp id")
		return
	}
	chirp, err := cfg.dbQueries.GetChirp(r.Context(), uuidId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Id not found")
		return
	}

	responseChirp := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, responseChirp)
}
