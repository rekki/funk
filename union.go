package funk

// Union returns the distinct values present in xs or ys. It preserves the
// order of each value's first occurrence, considering xs before ys.
func Union[T comparable](xs, ys []T) []T {
	output := make([]T, 0, len(xs)+len(ys))
	seen := make(map[T]struct{}, len(xs)+len(ys))
	for _, items := range [][]T{xs, ys} {
		for _, item := range items {
			if _, ok := seen[item]; ok {
				continue
			}
			seen[item] = struct{}{}
			output = append(output, item)
		}
	}
	return output
}
