package main

import (
	"MODULE_PATH/internal/auth"
	"MODULE_PATH/internal/database"
	"encoding/json"
	"net/http"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	var parm createUserRequest
	decoer := json.NewDecoder(r.Body)
	err := decoer.Decode(&parm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	cnx := r.Context()
	hashedPassword, err := auth.HashPassword(parm.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid password")
		return
	}

	user, err := cfg.dbQueries.CreateUser(cnx, database.CreateUserParams{
		Email:          parm.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	responseUser := User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}
	respondWithJSON(w, http.StatusCreated, responseUser)
}
