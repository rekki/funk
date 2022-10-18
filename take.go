package funk

func Take[A any](xs []A, count int) []A {
	if count > len(xs) {
		return xs
	}
	return xs[:count]
}
