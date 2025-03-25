package funk

// Take returns the first n elements of a slice
func Take[A any](xs []A, count int) []A {
	if count > len(xs) {
		return xs
	}
	return xs[:count]
}

func TakeString(s string, count int) string {
	if count > len(s) {
		return s
	}
	return s[:count]
}
