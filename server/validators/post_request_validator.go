package validators

import (
	"html"
	"net/http"
	"strconv"
	"strings"
)

// validates a request for posts index.
// returns:
// - int: HTTP status code.
// - string: Error or success message.
// - int: category_id if the action type is fetch by 'category'.
func IndexPostsRequest(r *http.Request) (int, string, string, int, int) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", "", 0, 0
	}

	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, "Failed to parse form data", "", 0, 0
	}

	category := r.FormValue("category_id")
	created := r.FormValue("created")
	liked := r.FormValue("liked")

	page := 0
	pageStr := r.FormValue("page")
	if pageStr != "" {
		page, err = strconv.Atoi(r.FormValue("page"))
		if err != nil || page < 1 {
			return http.StatusBadRequest, "Invalid page number", "", 0, 0
		}
		page-- // in the databse the page number should start from 0
	}

	limit := 10
	limitStr := r.FormValue("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			return http.StatusBadRequest, "Invalid limit number", "", 0, 0
		}
	}
	if category == "" && created == "" && liked == "" {
		return http.StatusOK, "success", "index", 0, page
	} else if category != "" && created == "" && liked == "" {
		categoryID, err := strconv.Atoi(category)
		if err != nil || categoryID < 1 {
			return http.StatusBadRequest, "Invalid category ID", "", 0, page
		}
		return http.StatusOK, "success", "category", categoryID, page
	} else if created != "" && category == "" && liked == "" {
		return http.StatusOK, "success", "created", 0, page
	} else if liked != "" && category == "" && created == "" {
		return http.StatusOK, "success", "liked", 0, page
	} else {
		return http.StatusBadRequest, "Only one action type is allowed", "", 0, page
	}
}

// validates a request to show a specific post.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - int: post ID.
func ShowPostRequest(r *http.Request) (int, string, int) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", 0
	}

	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, "Failed to parse form data", 0
	}

	postIdStr := r.PathValue("id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil || postId < 1 {
		return http.StatusBadRequest, "Post ID must be a valid positive integer", 0
	}

	return http.StatusOK, "success", postId
}

// validates a request to create a new post.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - string: title of the post.
// - string: content of the post.
// - []int: List of category IDs.
func CreatePostRequest(r *http.Request) (int, string, string, string, []int) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", "", "", nil
	}

	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		return http.StatusUnsupportedMediaType, "Unsupported content type", "", "", nil
	}

	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, "Failed to parse form data", "", "", nil
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	if strings.TrimSpace(title) == "" {
		return http.StatusBadRequest, "Title is required", "", "", nil
	}
	if len(title) > 100 {
		return http.StatusBadRequest, "Title must not exceed 100 characters", "", "", nil
	}

	if len(categories) == 0 {
		return http.StatusBadRequest, "At least one category is required", "", "", nil
	}

	convertCategories := make([]int, 0, len(categories))
	for _, cat := range categories {
		if cat == "" {
			return http.StatusBadRequest, "Category ID cannot be empty", "", "", nil
		}

		categoryID, err := strconv.Atoi(cat)
		if err != nil {
			return http.StatusBadRequest, "Category ID must be a valid integer", "", "", nil
		}

		convertCategories = append(convertCategories, categoryID)
	}

	if strings.TrimSpace(content) == "" {
		return http.StatusBadRequest, "Content is required", "", "", nil
	}
	if len(content) > 3000 {
		return http.StatusBadRequest, "Content must not exceed 3000 characters", "", "", nil
	}

	return http.StatusOK, "success",
		html.EscapeString(title),
		html.EscapeString(content),
		convertCategories
}
