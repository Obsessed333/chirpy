package main

import (
	"net/http"
	"github.com/obsessed333/chirpy/internal/auth"
	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r * http.Request){
	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, 401, "Token not found")
		return
	}
	userID , err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil{
		respondWithError(w, 403, "Unauthorized")
		return
	}

	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil{
		respondWithError(w, 400, "Invalid chirp ID")
		return
	}
	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil{
		respondWithError(w, 404, "Chirp not found")
		return
	}
	if chirp.UserID != userID{
		respondWithError(w, 403, "Unauthorized")
		return
	}
	err = cfg.db.DeleteChirp(r.Context(), chirpID)
	if err != nil{
		respondWithError(w, 404, "Chirp not found")
		return
	}
	respondWithJSON(w, 204, "Chirp successfully deleted")
}