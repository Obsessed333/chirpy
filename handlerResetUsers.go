package main

import (
	"context"
	"net/http"
)


func (cfg *apiConfig) handlerResetUsers(w http.ResponseWriter, r *http.Request){
	if cfg.platform != "dev"{
		respondWithError(w, 403, "Forbidden")
		return
	}
	err := cfg.db.DeleteUsers(context.Background())
	if err != nil{
		respondWithError(w, 500, err.Error())
		return
	}
	respondWithJSON(w, 200, "Users table reset was successful")
}