package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error){
	apikey:= headers.Get("Authorization")
	if apikey==""{
		return "", errors.New("authorization header missing")
	}
	result := strings.TrimPrefix(apikey, "ApiKey ")
	return result, nil
}
