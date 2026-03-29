package main

import(
	"encoding/json"
	"net/http"
	"github.com/obsessed333/chirpy/internal/auth"
	"time"
	"github.com/obsessed333/chirpy/internal/database"
)


func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Password string `json:"password"`
		Email string `json:"email"`
	}
	
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	userdb, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil{
		respondWithError(w, 401, "Incorrect email or password")
		return
	}
	match, err := auth.CheckPasswordHash(params.Password, userdb.HashedPassword)
	if err != nil || !match{
		respondWithError(w, 401, "Incorrect email or password")
		return
	}
	user := User{
		ID: userdb.ID,
		CreatedAt: userdb.CreatedAt,
		UpdatedAt: userdb.UpdatedAt,
		Email: userdb.Email,
		IsChirpyRed: userdb.IsChirpyRed,
	}
	type response struct {
    User
    Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	}
	
	jwt, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(3600) * time.Second)
	if err != nil{
		respondWithError(w, 500, "Error making JWT")
		return
	}
	refreshToken := auth.MakeRefreshToken()
	refreshTokenDB, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token: refreshToken,
		UserID: user.ID,
		ExpiresAt: time.Now().UTC().Add(60 * 24 * time.Hour),
	})
	if err != nil{
		respondWithError(w, 500, "Error creating refresh token")
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			IsChirpyRed: user.IsChirpyRed,
		},
		Token: jwt,
		RefreshToken: refreshTokenDB.Token,
	})
}