package funk

// Union returns a new slice containing all distinct elements present in both xs and ys
func Union[T comparable](xs, ys []T) []T {
	output := make([]T, 0, len(xs)+len(ys))

	if xs == nil && ys == nil {
		return output
	}

	if xs != nil && ys == nil {
		output = append(output, xs...)
		return output
	}

	if ys != nil && xs == nil {
		output = append(output, ys...)
		return output
	}

	if len(xs) == 0 && len(ys) == 0 {
		return output
	}

	m := make(map[T]bool)

	for _, x := range xs {
		m[x] = true
		output = append(output, x)
	}

	for _, y := range ys {
		if _, ok := m[y]; !ok {
			output = append(output, y)
		}
	}
	return output
}
