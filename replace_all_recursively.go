package funk

import "strings"

// ReplaceAllRecursively replaces every occurrence of toReplace, repeating
// while each pass shortens the result. This is useful when a replacement can
// create a new match, such as collapsing repeated whitespace. An empty
// toReplace leaves text unchanged.
func ReplaceAllRecursively(text string, toReplace string, replaceWith string) string {
	if toReplace == "" {
		return text
	}

	current := text
	for {
		next := strings.ReplaceAll(current, toReplace, replaceWith)
		if len(next) >= len(current) {
			return next
		}
		current = next
	}
}
