package auth

import(
	"net/http"
	"errors"
	"strings"
)

func GetAPIKey(headers http.Header) (string,error){

	apiKey := headers.Get("Authorization")
	if apiKey == ""{
		return "", errors.New("Authorization header is empty")
	}
	parts := strings.Split(apiKey, " ")
	if parts[0] != "ApiKey" || len(parts) < 2{
		return "", errors.New("Authorization header is malformed")
	}
	return parts[1], nil
}