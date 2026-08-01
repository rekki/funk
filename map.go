package funk

// Map applies mapper to each element of items. It always returns a non-nil
// slice.
func Map[R, T any](items []T, mapper func(T) R) []R {
	out := make([]R, len(items))
	for index, item := range items {
		out[index] = mapper(item)
	}
	return out
}
