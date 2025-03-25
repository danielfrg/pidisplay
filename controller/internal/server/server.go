package server

import (
	"log"
	"encoding/json"
	"net/http"
	"strconv"
)

type APIHandler struct{}

// GetDisplays handles retrieving display information
func (h *APIHandler) GetDisplays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConfigData.Displays)
}

// GetScreenshot handles taking and returning a screenshot
func (h *APIHandler) GetScreenshot(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Get the required "url" query parameter
	urlParam := query.Get("url")
	if urlParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "url query parameter is required"})
		return
	}

	// Read optional parameters with defaults
	widthStr := query.Get("width")
	if widthStr == "" {
		widthStr = "1280"
	}

	heightStr := query.Get("height")
	if heightStr == "" {
		heightStr = "720"
	}

	binary := query.Get("binary")
	if binary == "" {
		binary = "chrome-headless-shell"
	}

	// Convert width and height to integers
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid width value"})
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid height value"})
		return
	}

	dims := [2]int{width, height}

	imageBytes, err := TakeScreenshot(urlParam, dims, binary)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageBytes)
}

// New returns a new http.Handler that serves the API endpoints.
func New() http.Handler {
	router := http.NewServeMux()

	// Get the static file handler
	assetsHandler, err := AssetsHandler()
	if err != nil {
		log.Fatalf("Error setting up assets: %v", err)

		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("PiDisplay Controller API"))
		})
	} else {
		// Serve the static files
		router.Handle("/", assetsHandler)
	}

	apiHandler := &APIHandler{}

	router.HandleFunc("GET /api/v1/screenshot", func(w http.ResponseWriter, r *http.Request) {
		apiHandler.GetScreenshot(w, r)
	})

	router.HandleFunc("GET /api/v1/displays", func(w http.ResponseWriter, r *http.Request) {
		apiHandler.GetDisplays(w, r)
	})

	return router
}
