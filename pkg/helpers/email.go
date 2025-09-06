package helpers

import "strings"

// NormalizeEmail normalizes the email by lowercasing and trimming spaces.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
