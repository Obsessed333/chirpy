package main

import (
	"net/http"
	"log"
	"os"
	"database/sql"
	"sync/atomic"
	"github.com/joho/godotenv"
  _ "github.com/lib/pq"
	"github.com/obsessed333/chirpy/internal/database"
)


func main (){
	
godotenv.Load()

dbURL := os.Getenv("DB_URL")
platform := os.Getenv("PLATFORM")
polkaKey := os.Getenv("POLKA_KEY")
jwtSecret := os.Getenv("JWT_SECRET")
db, err := sql.Open("postgres", dbURL)
if err != nil{
	log.Fatalf("ERROR OPENING DATABASE: %v", err)
}
dbQueries := database.New(db)

mux := http.NewServeMux()
handlerFileServer := http.FileServer(http.Dir("."))
mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
})

apiCfg := apiConfig{
	fileserverHits: atomic.Int32{},
	db: dbQueries,
	platform: platform,
	jwtSecret: jwtSecret,
	polkaKey: polkaKey,
}

mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", handlerFileServer)))
mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
mux.HandleFunc("POST /api/chirps", apiCfg.handlerCreateChirp)
mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
mux.HandleFunc("PUT /api/users", apiCfg.handlerUpdateUser)
mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetUsers)
mux.HandleFunc("GET /api/chirps/", apiCfg.handlerRetrieveChirps)
mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerGetChirp)
mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
mux.HandleFunc("POST /api/refresh", apiCfg.handlerLookUpRefreshToken)
mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevokeToken)
mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.handlerDeleteChirp)
mux.HandleFunc("POST /api/polka/webhooks", apiCfg.handlerUpgradeUser)

s := &http.Server{
	Addr: ":8080",
	Handler: mux,
}

if err := s.ListenAndServe(); err != nil{
	log.Fatal("Listen and serve error", err)
}
}

