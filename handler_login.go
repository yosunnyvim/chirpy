package main

import (
	"MODULE_PATH/internal/auth"
	"encoding/json"
	"net/http"
)
type loginUserRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}
func (cfg *apiConfig) login(w http.ResponseWriter, r *http.Request){
	var parm loginUserRequest
	decoer := json.NewDecoder(r.Body)
	err := decoer.Decode(&parm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	user,err:= cfg.dbQueries.GetUserByEmail(r.Context(),parm.Email)
	if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
			return
	}
	validLogin, err:= auth.CheckPasswordHash(parm.Password, user.HashedPassword)
	if err!= nil{
			respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
			return
	}
	if validLogin{
	responseData:= User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	}
	respondWithJSON(w, http.StatusOK, responseData)
	}else {
			respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
			return
	}
}
