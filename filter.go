package funk

// Filter a collection by a predicate.
func Filter[A any](items []A, predicate func(A) bool) []A {
	var out []A

	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}

	return out
}

// Filter a collection by a predicate, will never return nil
func FilterToEmpty[A any](items []A, predicate func(A) bool) []A {
	out := make([]A, 0, len(items))

	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}

	return out
}
