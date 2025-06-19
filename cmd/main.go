package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"forum/server/controllers"
	"forum/server/middlewares"
	"forum/server/models"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	if _, err := template.ParseFiles("./web/index.html"); err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	if err := models.Connect(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	if err := models.CreateTables(); err != nil {
		log.Fatalf("error creating demo data: %v", err)
	}

	// if err := models.CreateDemoData(); err != nil {
	// 	log.Fatalf("error creating demo data: %v", err)
	// }
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/api/", controllers.ServeAvailableRoutes)
	mux.HandleFunc("/api/assets/", controllers.ServeStaticFiles)
	mux.HandleFunc("/api/register", controllers.Register)
	mux.HandleFunc("/api/login", controllers.Login)

	// Routes that require authentication
	mux.HandleFunc("/api/users", middlewares.IsAuth(controllers.IndexUsers))
	mux.HandleFunc("/api/posts", middlewares.IsAuth(controllers.IndexPosts))
	mux.HandleFunc("/api/posts/{id}", middlewares.IsAuth(controllers.GetPostById))
	mux.HandleFunc("/api/posts/{id}/comments", middlewares.IsAuth(controllers.GetComments))
	mux.HandleFunc("/api/posts/create", middlewares.IsAuth(controllers.CreatePost))
	mux.HandleFunc("/api/posts/react", middlewares.IsAuth(controllers.ReactToPost))
	mux.HandleFunc("/api/comments/create", middlewares.IsAuth(controllers.CreateComment))
	mux.HandleFunc("/api/conversation/{sender}", middlewares.IsAuth(controllers.GetConvertation))
	mux.HandleFunc("/api/logout", middlewares.IsAuth(controllers.Logout))

	// WebSocket endpoint
	mux.HandleFunc("/ws", middlewares.IsAuth(controllers.HandleWebSocket))

	rateLimiter := middlewares.NewRateLimiter(1500, 1*time.Minute)

	return middlewares.Recovery(rateLimiter.Middleware(mux))
}

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Too many arguments")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: routes(),
	}

	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
