package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"forum/server/config"
	"forum/server/utils"
)

// ServeStaticFiles returns a handler function for serving static files
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Get clean file path and prevent directory traversal
	filePath := filepath.Clean(config.BasePath + "web" + r.URL.Path)

	// block access to dirictories
	if info, err := os.Stat(filePath); err != nil || info.IsDir() {
		utils.RenderError(nil, w, r, http.StatusNotFound, false, "")
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
