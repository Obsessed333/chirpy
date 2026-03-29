package main

import (
	"encoding/json"
	"net/http"
	"github.com/obsessed333/chirpy/internal/database"
	"github.com/obsessed333/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Body string `json:"body"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, 404, "Token not found")
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
	
	if len(params.Body) > 140{
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	cleanString := ProfaneFilter(params.Body)

	chirpdb, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body: cleanString,
		UserID: userID,
	})
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	chirp := Chirp{
		ID: chirpdb.ID,
		CreatedAt: chirpdb.CreatedAt,
		UpdatedAt: chirpdb.UpdatedAt,
		Body: chirpdb.Body,
		UserID: chirpdb.UserID,
	}
	respondWithJSON(w, 201, chirp)
}