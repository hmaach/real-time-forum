package api

import (
	"net/http"

	"forum/server/controllers"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	// render HTML index page
	mux.HandleFunc("/", controllers.Index)

	// Serve static files
	mux.HandleFunc("/assets/", controllers.ServeAssetsFiles)

	// API routes
	mux.HandleFunc("/api", controllers.IndexAPIs)
	mux.HandleFunc("/api/posts", controllers.IndexPosts)
	mux.HandleFunc("/api/posts/{id}", controllers.ShowPost)
	mux.HandleFunc("/api/posts/create", controllers.CreatePost)
	mux.HandleFunc("/api/posts/react", controllers.ReactToPost)
	mux.HandleFunc("/api/comment/create", controllers.CreateComment)
	mux.HandleFunc("/api/comment/react", controllers.ReactToComment)
	mux.HandleFunc("/api/register", controllers.Register)
	mux.HandleFunc("/api/login", controllers.Login)
	mux.HandleFunc("/api/logout", controllers.Logout)


return mux
}
