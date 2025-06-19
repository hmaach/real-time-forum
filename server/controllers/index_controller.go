package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"forum/server/models"
	"forum/server/utils"
)

// Index handles the root route and serves the index.html template
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Execute the pre-parsed template
	t, _ := template.ParseFiles("./web/index.html")
	if err := t.Execute(w, nil); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

// IndexUsers handles the root route and serves the index.html template
func IndexUsers(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Fetch users from the database and return them as JSON
	userID := r.Context().Value("user_id").(int)
	users, err := models.GetUsers(userID)
	if err != nil {
		log.Println("Failed to fetch users: ", err)
		utils.JSONResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"users": users, "status": http.StatusOK})
}

// ServeStaticFiles returns a handler function for serving static files
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	filePath := filepath.Clean("./web" + strings.TrimPrefix(r.URL.Path, "/api"))

	if info, err := os.Stat(filePath); err != nil || info.IsDir() {
		utils.JSONResponse(w, http.StatusNotFound, "Page not found")
		return
	}

	http.ServeFile(w, r, filePath)
}

func ServeAvailableRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	log.Println(r.URL.Path)
	if r.URL.Path != "/api/" {
		utils.JSONResponse(w, http.StatusNotFound, "Page not found")
		return
	}

	// Map of routes to descriptions
	routes := map[string]string{
		"/":                          "Index page",
		"/api":                       "List of available API routes",
		"/api/assets/":               "Serve static files",
		"/api/register":              "User registration endpoint",
		"/api/login":                 "User login endpoint",
		"/api/users":                 "Get a list of users (requires authentication)",
		"/api/posts":                 "Get a list of posts (requires authentication)",
		"/api/posts/{id}":            "Get a specific post by ID (requires authentication)",
		"/api/posts/{id}/comments":   "Get comments for a specific post (requires authentication)",
		"/api/posts/create":          "Create a new post (requires authentication)",
		"/api/posts/react":           "React to a post (requires authentication)",
		"/api/comments/create":       "Create a comment (requires authentication)",
		"/api/conversation/{sender}": "Get a conversation by sender (requires authentication)",
		"/api/logout":                "Log out the user (requires authentication)",
		"/ws":                        "WebSocket endpoint for real-time communication (requires authentication)",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(routes); err != nil {
		http.Error(w, "Failed to encode routes", http.StatusInternalServerError)
		return
	}
}
