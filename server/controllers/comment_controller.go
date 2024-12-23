package controllers

import (
	"database/sql"
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"strings"

	"forum/server/models"
	"forum/server/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Validate session
	userID, username, valid := models.ValidSession(r, db)

	// Validate method
	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content := html.EscapeString(strings.TrimSpace(r.FormValue("comment")))
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)

	if err != nil || strings.TrimSpace(content) == "" || len(strings.TrimSpace(content)) > 500 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Store the comment using the models package
	commentID, err := models.StoreComment(db, userID, postID, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Fetch additional details using the models package
	commentsCount, err := models.CountCommentsByPostID(db, postID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentTime, err := models.FetchCommentTimeByID(db, commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the new comment details as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ID":            commentID,
		"username":      username,
		"created_at":    commentTime,
		"content":       content,
		"likes":         0,
		"dislikes":      0,
		"commentscount": commentsCount,
	})
}

func ReactToComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	user_id, username, valid := models.ValidSession(r, db)

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	if !valid {
		w.WriteHeader(401)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("comment_id")
	comment_id, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	likeCount, dislikeCount, err := models.ReactToComment(db, user_id, comment_id, userReaction)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"commentlikesCount": likeCount, "commentdislikesCount": dislikeCount})
}
