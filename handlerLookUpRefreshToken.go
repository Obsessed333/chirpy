package main

import(
	"net/http"
	"time"
	"github.com/obsessed333/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLookUpRefreshToken(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Token string `json:"token"`
	}
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	refreshToken, err := cfg.db.GetUserFromRefreshToken(r.Context(), bearerToken)
	if err != nil{
		respondWithError(w, 401, "Something went wrong")
		return
	}
	jwt, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	respondWithJSON(w, 200, parameters{
		Token: jwt,
	})
}