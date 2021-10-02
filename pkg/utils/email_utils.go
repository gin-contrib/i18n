package utils

import (
	"regexp"
	"strings"
)

var (
	emailReCompile = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
)

// ValidateEmail --
func ValidateEmail(email string) bool {
	loweremail := strings.ToLower(email)
	return emailReCompile.MatchString(loweremail)
}
