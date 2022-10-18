package funk

func Some[A any](xs []A, predicate func(A) bool) bool {
	for _, x := range xs {
		if predicate(x) {
			return true
		}
	}
	return false
}
