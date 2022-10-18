package funk

//Filter a collection by a predicate.
func Filter[A comparable](from []A, filter func(A) bool) []A {
	var out []A

	for _, item := range from {
		if filter(item) {
			out = append(out, item)
		}
	}

	return out
}
