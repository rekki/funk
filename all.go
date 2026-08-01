package funk

// All reports whether predicate accepts every value in items.
// It reports true for an empty slice.
func All[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}
