package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"forum/server/config"
)

// serve the Index page
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles(config.BasePath + "web/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func IndexAPIs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.APIs)
}

// ServeStaticFiles returns a handler function for serving static files
func ServeAssetsFiles(w http.ResponseWriter, r *http.Request) {
	// Get clean file path and prevent directory traversal
	filePath := filepath.Clean(config.BasePath + "web" + r.URL.Path)

	// block access to dirictories and html files
	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() || filepath.Ext(filePath) == ".html" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
