package main

import (
	"MODULE_PATH/internal/auth"
	"net/http"
	"time"
)

type refreshTokenResponse struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) refresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unable to get Token")
		return
	}
	refreshToken, err := cfg.dbQueries.GetRefreshTokenByToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unable to get Token")
		return
	}
	if time.Now().UTC().After(refreshToken.ExpiresAt) {
		respondWithError(w, http.StatusUnauthorized, "Old token")
		return
	}
	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Already revoked")
		return
	}
	refreshedToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create access token")
		return
	}
	returned := refreshTokenResponse{
		Token: refreshedToken,
	}
	respondWithJSON(w, http.StatusOK, returned)
}
