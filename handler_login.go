package main

import (
	"MODULE_PATH/internal/auth"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)
type loginUserRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
	ExpiresInSeconds int `json:"expires_in_seconds"`
}
type ResponseLogin struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token     string    `json:"token"` 
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
	if !validLogin{
			respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
			return
	}	
	expiresTime:=parm.ExpiresInSeconds
	if expiresTime == 0{
		expiresTime =3600 
	}
	if parm.ExpiresInSeconds>3600{
		expiresTime=3600
	}
	duration:= time.Duration(expiresTime)*time.Second
	JWT,err:=auth.MakeJWT(user.ID, cfg.jwtSecret, duration)
	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return 
	}
	responseData:= ResponseLogin{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
		Token: JWT,	
	}
	respondWithJSON(w, http.StatusOK, responseData)

	}
