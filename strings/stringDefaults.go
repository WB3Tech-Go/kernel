package strings

import "strings"

// CleanString removes all leading & trailing spaces, or line breaks, in a string,
// while retaining any other spaces or line breaks.
func CleanString(value string) string {
	return strings.TrimSpace(value)
}