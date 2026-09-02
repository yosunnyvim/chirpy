package main

import (
	"encoding/json"
	"net/http"
)

type createUserRequest struct {
	Email string `json:"email"`
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
	user, err := cfg.dbQueries.CreateUser(cnx, parm.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	responseUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
	respondWithJSON(w, http.StatusCreated, responseUser)
}
