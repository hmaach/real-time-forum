package routes

import (
	"database/sql"
	"net/http"

	"forum/server/controllers"
)

func Routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// Serve static files
	mux.HandleFunc("/assets/", controllers.ServeStaticFiles)

	// routes to get pages
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPosts(w, r, db)
	})
	mux.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPostsByCategory(w, r, db)
	})
	mux.HandleFunc("/mycreatedposts", func(w http.ResponseWriter, r *http.Request) {
		controllers.MyCreatedPosts(w, r, db)
	})
	
	mux.HandleFunc("/mylikedposts", func(w http.ResponseWriter, r *http.Request) {
		controllers.MyLikedPosts(w, r, db)
	})
	mux.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.ShowPost(w, r, db)
	})

	mux.HandleFunc("/post/addcommentREQ", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateComment(w, r, db)
	})

	mux.HandleFunc("/post/create", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetPostCreationForm(w, r, db)
	})

	mux.HandleFunc("/post/createpost", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePost(w, r, db)
	})

	mux.HandleFunc("/post/postreaction", func(w http.ResponseWriter, r *http.Request) {
		controllers.ReactToPost(w, r, db)
	})

	mux.HandleFunc("/post/commentreaction", func(w http.ResponseWriter, r *http.Request) {
		controllers.ReactToComment(w, r, db)
	})

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		controllers.Signin(w, r, db)
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.Signup(w, r, db)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetLoginPage(w, r, db)
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.Logout(w, r, db)
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetRegisterPage(w, r, db)
	})

	return mux
}
