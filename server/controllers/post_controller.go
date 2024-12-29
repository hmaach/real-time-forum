package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"
)

func IndexPosts(w http.ResponseWriter, r *http.Request) {
	status, message, actionType, categoryID, page := validators.IndexPostsRequest(r)
	if status != http.StatusOK {
		utils.ErrorJSONResponse(w, status, message)
		return
	}
	// userID := 11

	var (
		posts []models.Post
		err   error
	)
	page *= 10

	switch actionType {
	case "index":
		posts, err = models.FetchPosts(page)
	case "category":
		posts, err = models.FetchPostsByCategory(categoryID, page)
		// case "created":
		// 	posts, err = models.FetchCreatedPostsByUser(userID)
		// case "liked":
		// 	posts, err = models.FetchLikedPostsByUser(userID)
	}

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	// Validate the request and extract the post ID
	status, message, postID := validators.ShowPostRequest(r)
	if status != http.StatusOK {
		utils.ErrorJSONResponse(w, status, message)
		return
	}

	post, err := models.FetchPost(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.ErrorJSONResponse(w, http.StatusNotFound, "Post not found")
		} else {
			log.Printf("Error fetching post with ID %d: %v", postID, err)
			utils.ErrorJSONResponse(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	status, _, title, content, categories := validators.CreatePostRequest(r)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	user_id, _, valid := models.ValidSession(r)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := models.CheckCategories(categories)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post_id, err := models.StorePost(user_id, title, content)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, category := range categories {
		_, err = models.StorePostCategory(post_id, category)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
}

func ReactToPost(w http.ResponseWriter, r *http.Request) {
	status, _, post_id, reactionType := validators.ReactRequest(r)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	user_id, _, valid := models.ValidSession(r)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	likeCount, dislikeCount, err := models.ReactToPost(user_id, post_id, reactionType)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"likesCount": likeCount, "dislikesCount": dislikeCount})
}
