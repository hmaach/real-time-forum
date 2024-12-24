package utils

import "regexp"

// IsValidEmail checks if the provided email is valid.
func IsValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
