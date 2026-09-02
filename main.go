package main

import (
	"net/http"
)

func main(){
	mux:= http.NewServeMux()
	server:= http.Server{
		Addr : ":8080",
		Handler: mux,
	}
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	apiCfg:= apiConfig{}
	handler:= http.StripPrefix("/app/",http.FileServer(http.Dir("./assets/")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.HandleFunc("GET /admin/metrics",apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.metricsRest)
	mux.HandleFunc("POST /api/validate_chirp",validateChrip)
	err:= server.ListenAndServe()
	if err!= nil {
		panic(err)
	}
}
