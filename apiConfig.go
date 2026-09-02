package main
import(
	"net/http"
	"sync/atomic"
	"fmt"
)
type apiConfig struct {
	fileserverHits atomic.Int32
}
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cfg.fileserverHits.Add(1)
			next.ServeHTTP(w,r)
	})	
}
func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	hits:=cfg.fileserverHits.Load()
	w.Write([]byte(fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %v times!</p></body></html>", hits)))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
func (cfg *apiConfig) metricsRest(w http.ResponseWriter, r *http.Request){
	cfg.fileserverHits.Store(0)
}
