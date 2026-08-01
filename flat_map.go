package funk

// FlatMap applies mapper to each element of items and concatenates the returned
// slices in input order.
func FlatMap[T, R any](items []T, mapper func(T) []R) []R {
	out := []R{}
	for _, item := range items {
		out = append(out, mapper(item)...)
	}
	return out
}
