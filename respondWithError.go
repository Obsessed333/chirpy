package main

import(
	"net/http"
)
func respondWithError(w http.ResponseWriter, code int, msg string){
		type errResp struct{
			Error string `json:"error"`
		}
		var jsonError errResp
		jsonError.Error = msg
		respondWithJSON(w, code, jsonError)
	}