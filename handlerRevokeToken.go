package main

import(
	"net/http"
	"github.com/obsessed333/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter,r *http.Request){
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	err = cfg.db.RevokeToken(r.Context(), bearerToken)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}