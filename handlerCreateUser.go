package main

import(
	"context"
	"github.com/google/uuid"
	"time"
	"encoding/json"
	"net/http"
	"github.com/obsessed333/chirpy/internal/database"
	"github.com/obsessed333/chirpy/internal/auth"
)



func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){

	type parameters struct {
    Email string `json:"email"`
	Password     string    `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}
	hashed, err := auth.HashPassword(params.Password)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	userdb, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email: params.Email,
		HashedPassword: hashed,
		IsChirpyRed: false,
	})
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	
	user := User{
		ID: userdb.ID,
		CreatedAt: userdb.CreatedAt,
		UpdatedAt: userdb.UpdatedAt,
		Email: userdb.Email,
		IsChirpyRed: userdb.IsChirpyRed,
	}
	respondWithJSON(w, 201, user)
}