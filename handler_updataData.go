package main

import (
	"MODULE_PATH/internal/auth"
	"MODULE_PATH/internal/database"
	"encoding/json"
	"net/http"
)

type updateDataRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}
type updateDataResponse struct {
    Email string    `json:"email"`
}
func (cfg *apiConfig)updateData(w http.ResponseWriter, r *http.Request){
	
	var parm updateDataRequest 
	decoer := json.NewDecoder(r.Body)
	err := decoer.Decode(&parm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	token,err:=auth.GetBearerToken(r.Header)
	if err!=nil{
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	valid,err:=auth.ValidateJWT(token, cfg.jwtSecret)
	if err!=nil{
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	hashedPass,err:=auth.HashPassword(parm.Password)
	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Invalid password")
		return
	}
	updated,err:=cfg.dbQueries.UpdateUser(r.Context(),database.UpdateUserParams{
		ID: valid,
		Email: parm.Email,
		HashedPassword: hashedPass,
	})	
	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}	
	returned:= updateDataResponse{
		Email: updated.Email,
	}

	respondWithJSON(w, http.StatusOK, returned)
	
}
