package funk

// Diff returns elements of xs that are not present in ys, preserving order.
// It returns nil when there are no such elements.
func Diff[A comparable](xs []A, ys []A) []A {
	ySet := make(map[A]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}
	var diff []A
	for _, x := range xs {
		if _, exists := ySet[x]; !exists {
			diff = append(diff, x)
		}
	}
	return diff
}

// Remove is like Diff but returns a non-nil empty slice when there are no
// remaining elements.
func Remove[A comparable](xs []A, ys []A) []A {
	ySet := make(map[A]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}
	diff := make([]A, 0)
	for _, x := range xs {
		if _, exists := ySet[x]; !exists {
			diff = append(diff, x)
		}
	}
	return diff
}

// Intersect returns elements of xs that are also present in ys, preserving
// order. It returns nil when there are no such elements.
func Intersect[A comparable](xs []A, ys []A) []A {
	ySet := make(map[A]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}
	var intersect []A
	for _, x := range xs {
		if _, exists := ySet[x]; exists {
			intersect = append(intersect, x)
		}
	}
	return intersect
}

// IntersectKeyMap returns a new map containing entries from m1 whose keys are
// also present in m2. Values from m2 are ignored.
func IntersectKeyMap[T any, K comparable](m1 map[K]T, m2 map[K]T) map[K]T {
	intersect := map[K]T{}

	for k, v := range m1 {
		if _, exist := m2[k]; exist {
			intersect[k] = v
		}
	}

	return intersect
}

// IdenticalSets reports whether xs and ys contain the same distinct values.
// Order and duplicate occurrences are ignored.
func IdenticalSets[A comparable](xs []A, ys []A) bool {
	xSet := make(map[A]struct{}, len(xs))
	for _, x := range xs {
		xSet[x] = struct{}{}
	}
	ySet := make(map[A]struct{}, len(ys))
	for _, y := range ys {
		ySet[y] = struct{}{}
	}
	if len(xSet) != len(ySet) {
		return false
	}
	for x := range xSet {
		if _, ok := ySet[x]; !ok {
			return false
		}
	}
	return true
}
