package funk

// UniqueBy returns the first item for each distinct derived key, preserving
// input order.
func UniqueBy[A any, B comparable](items []A, getKey func(A) B) []A {
	result := make([]A, 0)
	seen := make(map[B]struct{})
	for _, item := range items {
		itemKey := getKey(item)
		if _, ok := seen[itemKey]; ok {
			continue
		}
		seen[itemKey] = struct{}{}
		result = append(result, item)
	}
	return result
}
