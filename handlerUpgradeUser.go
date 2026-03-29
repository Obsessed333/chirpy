package main

import(
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/obsessed333/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter,r *http.Request){
	type parameters struct {
		Event string `json:"event"`
		Data  struct{
			UserID uuid.UUID `json:"user_id"`
		}
	}
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil{
		respondWithError(w, 401, "Unauthorized")
		return
	}
	if cfg.polkaKey != apiKey{
		respondWithError(w, 401, "Unauthorized")
		return
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	if params.Event != "user.upgraded"{
		w.WriteHeader(http.StatusNoContent)
		return
	}


	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), params.Data.UserID)
	if err != nil{
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}