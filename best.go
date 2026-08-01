package funk

// Best selects an item by repeatedly applying replace. The replace function
// reports whether next should replace current. Best returns false for an empty
// slice.
func Best[T any](items []T, replace func(current, next T) bool) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}

	best := items[0]
	for _, next := range items[1:] {
		if replace(best, next) {
			best = next
		}
	}
	return best, true
}
