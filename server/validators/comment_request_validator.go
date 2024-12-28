package validators

import (
	"net/http"
	"strconv"
	"strings"
)

// Validates a request to create a comment.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - string: Comment content.
// - int: Post ID.
func CreateCommentRequest(r *http.Request) (int, string, string, int) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", "", 0
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, "Failed to parse form data", "", 0
	}

	// Validate comment content
	content := strings.TrimSpace(r.FormValue("comment"))
	if content == "" {
		return http.StatusBadRequest, "Comment content cannot be empty", "", 0
	}
	if len(content) > 1800 {
		return http.StatusBadRequest, "Comment content exceeds the maximum length of 1800 characters", "", 0
	}

	// Validate Post ID
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil || postID <= 0 {
		return http.StatusBadRequest, "Post ID must be a valid positive integer", "", 0
	}

	return http.StatusOK, "success", content, postID
}
