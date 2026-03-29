package auth

import (
	"net/http"
	"errors"
	"strings"
)


func GetBearerToken(headers http.Header) (string, error){
	bearerToken := headers.Get("Authorization")
	if bearerToken == ""{
		return "", errors.New("Authorization header is empty")
	}
	parts := strings.Split(bearerToken, " ")
	if parts[0] != "Bearer" || len(parts) < 2{
		return "", errors.New("Authorization header is malformed")
	}
	return parts[1], nil
}