package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error){
	var registerClaim jwt.RegisteredClaims
	token,err:= jwt.ParseWithClaims(
	tokenString,
	&registerClaim,
	func(token *jwt.Token) (any, error) {
			return []byte(tokenSecret), nil
	},
 ) 
	if err!=nil{
		return uuid.UUID{},err
	}
	if !token.Valid{
		return uuid.UUID{}, fmt.Errorf("invalid token")
	}
	userID, err := uuid.Parse(registerClaim.Subject)
	if err!=nil{
		return uuid.UUID{},err
	}
	return userID,nil
}
