package funk

// Unique returns a new slice containing each distinct value from items once,
// preserving first-occurrence order.
func Unique[A comparable](items []A) []A {
	result := make([]A, 0)
	seen := make(map[A]struct{})
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}
