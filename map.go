package funk

// Map applies mapper to each element of items.
func Map[T, R any](items []T, mapper func(T) R) []R {
	out := make([]R, len(items))
	for index, item := range items {
		out[index] = mapper(item)
	}
	return out
}
