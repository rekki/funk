package funk

// Range returns the subslice xs[start:end], with out-of-range bounds clamped
// to valid indexes. A negative end is relative to len(xs), so -1 excludes the
// final element. If the normalized start is not before end, Range returns an
// empty slice. The result shares storage with xs.
func Range[A any](xs []A, start int, end int) []A {
	if len(xs) == 0 {
		return []A{}
	}
	if start < 0 {
		start = 0
	}
	if start >= len(xs) {
		return []A{}
	}
	if end < 0 {
		end += len(xs)
	}
	if end < 0 {
		end = 0
	}
	if end > len(xs) {
		end = len(xs)
	}
	if start >= end {
		return []A{}
	}
	return xs[start:end]
}
