package main

import (
	"net/http"
	"pidisplay/server"
)

// corsMiddleware sets CORS headers for incoming requests.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set standard CORS headers.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// For preflight requests, respond with OK status.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := server.New()

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: corsMiddleware(router),
	}

	server.ListenAndServe()
}
