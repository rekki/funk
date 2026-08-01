package funk

// Every reports whether all elements in xs satisfy predicate. It reports true
// for an empty slice.
func Every[A any](xs []A, predicate func(A) bool) bool {
	for _, x := range xs {
		if !predicate(x) {
			return false
		}
	}
	return true
}
