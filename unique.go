package funk

// Unique returns each distinct value from items once, preserving first-
// occurrence order. It always returns a non-nil slice.
func Unique[T comparable](items []T) []T {
	result := make([]T, 0, len(items))
	seen := make(map[T]struct{}, len(items))
	for _, item := range items {
		if _, exists := seen[item]; exists {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}
