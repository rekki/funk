package funk

import "fmt"

// PrintMap formats each key-value pair for debugging. The order of the
// returned strings is unspecified.
func PrintMap[T any, K comparable](input map[K]T) []string {
	out := make([]string, 0, len(input))
	for key, value := range input {
		out = append(out, fmt.Sprintf("key: %v; value: %v", key, value))
	}
	return out
}
