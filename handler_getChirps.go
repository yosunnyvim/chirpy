package main

import (
	"MODULE_PATH/internal/database"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirps(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")

	var chirps []database.Chirp
	var err error

	if authorID == "" {
		chirps, err = cfg.dbQueries.GetChirps(r.Context())
	} else {
		id, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID")
			return
		}

		chirps, err = cfg.dbQueries.GetChirpsByAuthor(r.Context(), id)
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}

	responseChirps := make([]Chirp, 0, len(chirps))
	for _, chirp := range chirps {
		responseChirps = append(responseChirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	sortOrder := r.URL.Query().Get("sort")

	if sortOrder == "desc" {
		sort.Slice(responseChirps, func(i, j int) bool {
			return responseChirps[i].CreatedAt.After(responseChirps[j].CreatedAt)
		})
	} else {
		sort.Slice(responseChirps, func(i, j int) bool {
			return responseChirps[i].CreatedAt.Before(responseChirps[j].CreatedAt)
		})
	}
	respondWithJSON(w, http.StatusOK, responseChirps)
}
