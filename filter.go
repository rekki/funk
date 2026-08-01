package funk

// Filter returns elements accepted by predicate, preserving their order.
func Filter[A any](items []A, predicate func(A) bool) []A {
	var out []A

	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}

	return out
}

// FilterToEmpty is like Filter but always returns a non-nil slice.
func FilterToEmpty[A any](items []A, predicate func(A) bool) []A {
	out := make([]A, 0, len(items))

	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}

	return out
}
