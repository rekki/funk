package funk

// Difference returns elements of xs that are not present in ys, preserving
// order and duplicate occurrences. It returns nil when there are no results.
func Difference[T comparable](xs, ys []T) []T {
	ySet := make(map[T]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}

	var difference []T
	for _, x := range xs {
		if _, exists := ySet[x]; !exists {
			difference = append(difference, x)
		}
	}
	return difference
}

// Intersect returns elements of xs that are also present in ys, preserving
// order and duplicate occurrences. It returns nil when there are no results.
func Intersect[T comparable](xs, ys []T) []T {
	ySet := make(map[T]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}

	var intersection []T
	for _, x := range xs {
		if _, exists := ySet[x]; exists {
			intersection = append(intersection, x)
		}
	}
	return intersection
}

// SetEqual reports whether xs and ys contain the same distinct values.
// Order and duplicate occurrences are ignored.
func SetEqual[T comparable](xs, ys []T) bool {
	matched := make(map[T]bool, len(xs))
	for _, x := range xs {
		matched[x] = false
	}

	remaining := len(matched)
	for _, y := range ys {
		seen, exists := matched[y]
		if !exists {
			return false
		}
		if !seen {
			matched[y] = true
			remaining--
		}
	}
	return remaining == 0
}
