package dom

import "strings"

// IsNotEmpty returns true is a string contains a non-whitespace character, else false
func IsNotEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}
