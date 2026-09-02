package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	hits := cfg.fileserverHits.Load()
	w.Write([]byte(fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %v times!</p></body></html>", hits)))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
