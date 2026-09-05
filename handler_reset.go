package main

import "net/http"

func (cfg *apiConfig) metricsRest(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	cnx := r.Context()
	err := cfg.dbQueries.DeleteUsers(cnx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete users")
		return
	}
	w.WriteHeader(http.StatusOK)
}
