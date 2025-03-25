package main

import (
	"flag"
	"net/http"
	"os"
	"pidisplay/internal/server"
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
	// Set default values from env if available.
	defaultHost := "localhost"
	defaultPort := "3000"

	if envHost := os.Getenv("PIDISPLAY_HOST"); envHost != "" {
		defaultHost = envHost
	}
	if envPort := os.Getenv("PIDISPLAY_PORT"); envPort != "" {
		defaultPort = envPort
	}

	// Define CLI flags with the env values as defaults.
	host := flag.String("host", defaultHost, "server host")
	port := flag.String("port", defaultPort, "server port")
	flag.Parse()

	// Construct the server address.
	addr := *host + ":" + *port

	router := server.New()

	srv := &http.Server{
		Addr:    addr,
		Handler: corsMiddleware(router),
	}

	// Start the server.
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
