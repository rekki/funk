package funk

// OneP returns an item in the list that is the most extreme in some way.
func OneP[A comparable](from []A, replace func(A, A) bool) A {
	if len(from) < 1 {
		panic("Cannot perform OneP on an empty array")
	}
	var max = from[0]
	for _, next := range from {
		if replace(max, next) {
			max = next
		}
	}
	return max
}
