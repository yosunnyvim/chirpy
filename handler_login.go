package main

import (
	"MODULE_PATH/internal/auth"
	"MODULE_PATH/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ResponseLogin struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	IsChirpyRed  bool      `json:"is_chirpy_red"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

func (cfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {
	var parm loginUserRequest
	decoer := json.NewDecoder(r.Body)
	err := decoer.Decode(&parm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	user, err := cfg.dbQueries.GetUserByEmail(r.Context(), parm.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	validLogin, err := auth.CheckPasswordHash(parm.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	if !validLogin {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	duration := time.Hour
	JWT, err := auth.MakeJWT(user.ID, cfg.jwtSecret, duration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return
	}
	refreshToken := auth.MakeRefreshToken()
	expiresAt := time.Now().Add(60 * 24 * time.Hour)
	_, err = cfg.dbQueries.CreateRefreshToken(
		r.Context(),
		database.CreateRefreshTokenParams{
			Token:     refreshToken,
			UserID:    user.ID,
			ExpiresAt: expiresAt,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create refreshToken")
		return
	}
	responseData := ResponseLogin{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		IsChirpyRed:  user.IsChirpyRed,
		Token:        JWT,
		RefreshToken: refreshToken,
	}
	respondWithJSON(w, http.StatusOK, responseData)

}
