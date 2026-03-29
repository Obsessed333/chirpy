package main

import(
	"net/http"
	"encoding/json"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
		jsonData, err := json.Marshal(payload)
		if err != nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonData)
}