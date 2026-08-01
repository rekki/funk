package funk

// UniqueBy returns the first item for each distinct derived key, preserving
// input order. It always returns a non-nil slice.
func UniqueBy[T any, K comparable](items []T, key func(T) K) []T {
	result := make([]T, 0, len(items))
	seen := make(map[K]struct{}, len(items))
	for _, item := range items {
		itemKey := key(item)
		if _, exists := seen[itemKey]; exists {
			continue
		}
		seen[itemKey] = struct{}{}
		result = append(result, item)
	}
	return result
}
