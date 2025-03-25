package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed "dist/**"
var Files embed.FS

// Returns an HTTP handler for the embedded static files
func AssetsHandler() (http.Handler, error) {
	subFS, err := fs.Sub(Files, "dist")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}
