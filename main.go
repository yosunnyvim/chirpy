package main

import (
	"database/sql"
	"net/http"
	"os"

	"MODULE_PATH/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	platform := os.Getenv("PLATFORM")
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	apiCfg := apiConfig{
		dbQueries: dbQueries,
		platform: platform,
	}
	handler := http.StripPrefix("/app/", http.FileServer(http.Dir("./assets/")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.metricsRest)
	mux.HandleFunc("POST /api/chirps", apiCfg.createChirp)
	mux.HandleFunc("POST /api/users", apiCfg.createUser)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
