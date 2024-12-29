package utils

import "unicode"

// Helper function to check if a string is alphanumeric
func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// Helper function to check if a string contains at least one uppercase letter
func ContainsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// Helper function to check if a string contains at least one digit
func ContainsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}