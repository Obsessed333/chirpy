package main

import(
	"net/http"
	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request){
	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil{
		respondWithError(w, 400, "Invalid chirp ID")
		return
	}
	chirpdb, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil{
		respondWithError(w, 404, "Chirp not found")
		return
	}
	chirp := Chirp{
		ID: chirpdb.ID,
		CreatedAt: chirpdb.CreatedAt,
		UpdatedAt: chirpdb.UpdatedAt,
		Body: chirpdb.Body,
		UserID: chirpdb.UserID,
	}
	respondWithJSON(w, 200, chirp)
}