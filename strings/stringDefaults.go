package strings

import "strings"

// cleanString removes all leading & trailing spaces, or line breaks, in a string,
// while retaining any other spaces or line breaks.
func cleanString(value string) string {
	return strings.TrimSpace(value)
}