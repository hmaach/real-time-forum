package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
)

func GenerateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// IsValidEmail checks if the provided email is valid.
func IsValidEmail(email string) bool {
	
	const emailRegex = `^[a-z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
