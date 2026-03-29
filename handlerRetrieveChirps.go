package main

import(
	"net/http"
	"github.com/google/uuid"
	"sort"
)


func (cfg *apiConfig) handlerRetrieveChirps(w http.ResponseWriter, r *http.Request){
	s := r.URL.Query().Get("author_id")
	sortQuery := r.URL.Query().Get("sort")
	if s != ""{
		authorID, err := uuid.Parse(s)
		if err != nil{
			respondWithError(w, http.StatusBadRequest, "Invalid author ID")
			return
		}
		authorChirps, err := cfg.db.GetChirpsByAuthor(r.Context(), authorID)
		if err != nil{
			respondWithError(w, 500, "Something went wrong")
			return
		}
		chirps := []Chirp{}
		for _, chirp := range authorChirps{
			chirps = append(chirps,Chirp{
				ID: chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body: chirp.Body,
				UserID: chirp.UserID,
		})
		}
		if sortQuery == "desc"{
			sort.Slice(chirps, func (i, j int) bool{
				return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
			})
		} else{
			sort.Slice(chirps, func (i, j int) bool{
			return chirps[j].CreatedAt.After(chirps[i].CreatedAt)
			})
		}
		respondWithJSON(w, 200, chirps)
		
	} else{
	dbChirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil{
		respondWithError(w, 500, "Something went wrong")
		return
	}
	apiChirps := []Chirp{}
	for _, chirp := range dbChirps{
			apiChirps = append(apiChirps,Chirp{
			ID: chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body: chirp.Body,
			UserID: chirp.UserID,
		})
	}
	if sortQuery == "desc"{
			sort.Slice(apiChirps, func (i, j int) bool{
				return apiChirps[i].CreatedAt.After(apiChirps[j].CreatedAt)
			})
		} else{
			sort.Slice(apiChirps, func (i, j int) bool{
			return apiChirps[j].CreatedAt.After(apiChirps[i].CreatedAt)
			})
		}
		respondWithJSON(w, 200, apiChirps)
}
}