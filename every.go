package funk

func Every[A any](xs []A, p func(A) bool) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}
	return true
}
