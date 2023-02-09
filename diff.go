package funk

// Diff returns all xs that are not in ys
func Diff[A comparable](xs []A, ys []A) []A {
	var diff []A = nil
	yMap := ToMap(ys, func(y A) A { return y })
	for _, x := range xs {
		_, exists := yMap[x]
		if !exists {
			diff = append(diff, x)
		}
	}

	return diff
}

// Intersect returns all xs that are also in ys
func Intersect[A comparable](xs []A, ys []A) []A {
	var intersect []A = nil
	yMap := ToMap(ys, func(y A) A { return y })
	for _, x := range xs {
		_, exists := yMap[x]
		if exists {
			intersect = append(intersect, x)
		}
	}

	return intersect
}

//IdenticalSets tests whether the size and elements of two sets are identical. Order does not matter
func IdenticalSets[A comparable](xs []A, ys []A) bool {
	if len(xs) > len(ys) {
		return false
	}
	if len(Diff(xs, ys)) > 0 {
		return false
	}
	if len(Diff(ys, xs)) > 0 {
		return false
	}
	return true
}
