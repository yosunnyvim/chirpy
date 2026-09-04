package main

import ( "database/sql"
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
	JWTSecret:= os.Getenv("JWT_SECRET")
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	apiCfg := apiConfig{
		dbQueries: dbQueries,
		platform:  platform,
		jwtSecret: JWTSecret,
	}
	mux.HandleFunc("GET /api/healthz", apiCfg.healthz )
	handler := http.StripPrefix("/app/", http.FileServer(http.Dir("./assets/")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.metricsRest)
	mux.HandleFunc("POST /api/chirps", apiCfg.createChirp)
	mux.HandleFunc("POST /api/users", apiCfg.createUser)
	mux.HandleFunc("GET /api/chirps", apiCfg.getChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.getChirp)
	mux.HandleFunc("POST /api/login", apiCfg.login)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
