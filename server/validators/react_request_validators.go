package validators

import (
	"net/http"
	"strconv"
	"strings"
)

// validates a request to react to a post.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - int: target ID.
// - string: reaction type (like/dislike).
func ReactRequest(r *http.Request) (int, string, int, string) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", 0, ""
	}

	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, "Failed to parse form data", 0, ""
	}

	// Validate the reaction type
	reactionType := strings.TrimSpace(r.FormValue("reaction"))
	if reactionType != "like" && reactionType != "dislike" {
		return http.StatusBadRequest, "Invalid reaction type", 0, ""
	}

	// Validate the target ID
	targetIdStr := r.FormValue("target_id")
	targetId, err := strconv.Atoi(targetIdStr)
	if err != nil {
		return http.StatusBadRequest, "Target ID must be a valid integer", 0, ""
	}

	if targetId <= 0 {
		return http.StatusBadRequest, "Target ID must be greater than 0", 0, ""
	}

	return http.StatusOK, "success", targetId, reactionType
}
