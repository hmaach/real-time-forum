package validators

import (
	"html"
	"net/http"
	"strings"
)

// Validates a login request.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - string: username.
// - string: password.
func LoginRequest(r *http.Request) (int, string, string, string) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", "", ""
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, "Failed to parse form data", "", ""
	}

	// Retrieve and sanitize inputs
	username := strings.TrimSpace(html.EscapeString(r.FormValue("username")))
	password := strings.TrimSpace(html.EscapeString(r.FormValue("password")))

	// Validate inputs
	if len(username) < 4 {
		return http.StatusBadRequest, "Username must be at least 4 characters long", "", ""
	}
	if len(password) < 6 {
		return http.StatusBadRequest, "Password must be at least 6 characters long", "", ""
	}

	return http.StatusOK, "success", username, password
}
