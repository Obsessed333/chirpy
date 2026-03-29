package auth

import(
	"testing"
	"github.com/google/uuid"
	"time"
	"net/http"
)
func TestJWT(t *testing.T){

	newUUID := uuid.New()
	secret := "secret-pw"
	token, err := MakeJWT(newUUID, secret, 50*time.Second)
	if err != nil{
		t.Errorf("error generating token: %v", err)
	}
	validatedID, err := ValidateJWT(token, secret)
	if err != nil{
		t.Errorf("error validating token: %v", err)
	}
	if validatedID != newUUID{
		t.Errorf("got %v, want %v", validatedID, newUUID)
	}
}

func TestPastExpirationDate(t *testing.T){
	newUUID := uuid.New()
	secret := "secret-pw"
	token , err := MakeJWT(newUUID, secret, -10*time.Second)
	if err != nil{
		t.Errorf("Error creating token: %v", err)
	}
	_ , err = ValidateJWT(token, "cuck")
	if err == nil{
		t.Errorf("Expected an error, but got nil")
	}
}

func TestWrongSecret(t *testing.T){
	newUUID := uuid.New()
	secret := "secret-pw"
	token, err := MakeJWT(newUUID, secret, 50*time.Second)
	if err != nil{
		t.Errorf("error generating token: %v", err)
	}
	_ , err = ValidateJWT(token, "cuck")
	if err == nil{
		t.Errorf("Expected an error, but got nil")
	}
}

func TestWhiteSpaceBearer(t *testing.T){
	header := http.Header{
		"Authorization": []string{"Bearer     extra-space-token"},
	}
	_ , err := GetBearerToken(header)
	if err != nil{
		t.Errorf("Error getting bearer token: %v", err)
	}
}