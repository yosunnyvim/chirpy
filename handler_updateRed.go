package main

import (
	"MODULE_PATH/internal/auth"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type polkaWebhookRequest struct {
	Event string `json:"event"`
	Data  struct {
		UserID uuid.UUID `json:"user_id"`
	} `json:"data"`
}

func (cfg *apiConfig) updateRed(w http.ResponseWriter, r *http.Request) {
	apiKey,err:=auth.GetAPIKey(r.Header)
	if err!=nil{
		respondWithError(w, http.StatusUnauthorized, "Invalid Request")
		return
	}
	if apiKey!=cfg.POLKA_KEY{
		respondWithError(w, http.StatusUnauthorized, "Invalid Request")
		return
	}
	var parm polkaWebhookRequest

	err = json.NewDecoder(r.Body).Decode(&parm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if parm.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.dbQueries.UpgradeUserToChirpyRed(
		r.Context(),
		parm.Data.UserID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
