package main

import (
	"time"
	"encoding/json"
	"net/http"
	"github.com/obsessed333/chirpy/internal/database"
	"github.com/obsessed333/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Password string `json:"password"`
		Email string `json:"email"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, 401, "Token not found")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil{
		respondWithError(w, 401, "Unauthorized")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err = decoder.Decode(&params)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}

	hashed, err := auth.HashPassword(params.Password)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	userDB, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		Email: params.Email, 
		HashedPassword: hashed,
		UpdatedAt: time.Now().UTC(),
		ID: userID,
	})
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	respondWithJSON(w, 200, User{
		ID: userDB.ID,
		CreatedAt: userDB.CreatedAt,
		Email: userDB.Email,
		UpdatedAt: userDB.UpdatedAt,
	})
}