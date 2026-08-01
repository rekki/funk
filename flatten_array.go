package funk

import (
	"encoding/json"
	"strings"
)

// FlattenArray formats items as a JSON-compatible array of quoted strings.
// Nil and empty slices are formatted as [].
func FlattenArray(items []string) string {
	quoted := Map(items, func(item string) string {
		encoded, _ := json.Marshal(item) // Marshaling a string cannot fail.
		return string(encoded)
	})
	return "[" + strings.Join(quoted, ", ") + "]"
}
