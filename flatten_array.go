package funk

import "strings"

// FlattenArray takes  a collection of strings (a,b,c) and returns an array of this form:
// [ "a", "b", "c" ]

func FlattenArray(items []string) string {

	if len(items) == 0 {
		return "[]"
	}

	withQuotes := Map(items, func(item string) string {
		return "\"" + item + "\""
	})

	withQuotesString := strings.Join(withQuotes, ", ")
	return "[" + withQuotesString + "]"

}
