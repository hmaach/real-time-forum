package validators

import (
	"html"
	"net/http"
	"strings"

	"forum/server/utils"
)

// Validates a registration request.
// Returns:
// - int: HTTP status code.
// - string: Error or success message.
// - string: email.
// - string: username.
// - string: password.
func RegisterRequest(r *http.Request) (int, string, string, string, string) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "Invalid HTTP method", "", "", ""
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, "Failed to parse form data", "", "", ""
	}

	// Retrieve and sanitize inputs
	email := strings.TrimSpace(html.EscapeString(r.FormValue("email")))
	username := strings.TrimSpace(html.EscapeString(r.FormValue("username")))
	password := strings.TrimSpace(r.FormValue("password"))
	passwordConfirmation := strings.TrimSpace(r.FormValue("password-confirmation"))

	// Validate email
	if email == "" {
		return http.StatusBadRequest, "Email is required", "", "", ""
	}
	if !utils.IsValidEmail(email) {
		return http.StatusBadRequest, "Invalid email format", "", "", ""
	}

	// Validate username
	if len(username) < 4 {
		return http.StatusBadRequest, "Username must be at least 4 characters long", "", "", ""
	}
	if strings.Contains(username, " ") {
		return http.StatusBadRequest, "Username cannot contain spaces", "", "", ""
	}
	if !utils.IsAlphanumeric(username) {
		return http.StatusBadRequest, "Username must contain only letters and numbers", "", "", ""
	}

	// Validate password
	if password != passwordConfirmation {
		return http.StatusBadRequest, "Passwords do not match", "", "", ""
	}
	if len(password) < 6 {
		return http.StatusBadRequest, "Password must be at least 6 characters long", "", "", ""
	}
	if !utils.ContainsUppercase(password) {
		return http.StatusBadRequest, "Password must contain at least one uppercase letter", "", "", ""
	}
	if !utils.ContainsDigit(password) {
		return http.StatusBadRequest, "Password must contain at least one digit", "", "", ""
	}

	return http.StatusOK, "success", email, username, password
}
