package controllers

import (
	"encoding/json"
	"net/http"

	"forum/server/models"
	"forum/server/validators"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	statusCode, _, content, postID := validators.CreateCommentRequest(r)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	// Validate session
	userID, _, valid := models.ValidSession(r)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Store the comment using the models package
	_, err := models.StoreComment(userID, postID, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReactToComment(w http.ResponseWriter, r *http.Request) {
	status, _, comment_id, reactionType := validators.ReactRequest(r)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	user_id, _, valid := models.ValidSession(r)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	likeCount, dislikeCount, err := models.ReactToComment(user_id, comment_id, reactionType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"commentlikesCount": likeCount, "commentdislikesCount": dislikeCount})
}
