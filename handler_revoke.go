package main

import (
	"MODULE_PATH/internal/auth"
	"net/http"
)

func (cfg *apiConfig) revoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unable to get token")
		return
	}
	err=cfg.dbQueries.RevokeRefreshToken(r.Context(), token)
	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Unable to revoke token")
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
