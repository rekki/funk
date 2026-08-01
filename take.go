package funk

// Take returns a subslice containing at most the first count elements of xs.
// A non-positive count returns an empty subslice. The result shares storage
// with xs.
func Take[A any](xs []A, count int) []A {
	if count <= 0 {
		return xs[:0]
	}
	if count >= len(xs) {
		return xs
	}
	return xs[:count]
}

// TakeString returns at most the first count Unicode code points in s. A
// non-positive count returns an empty string.
func TakeString(s string, count int) string {
	if count <= 0 {
		return ""
	}

	runes := 0
	for i := range s {
		if runes == count {
			return s[:i]
		}
		runes++
	}
	return s
}
