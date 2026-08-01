package funk

// Filter returns elements accepted by predicate, preserving their order.
func Filter[T any](items []T, predicate func(T) bool) []T {
	out := make([]T, 0, len(items))
	for _, item := range items {
		if predicate(item) {
			out = append(out, item)
		}
	}
	return out
}

// FilterMap maps the values accepted by fn, preserving their order.
func FilterMap[T, R any](items []T, fn func(T) (R, bool)) []R {
	out := make([]R, 0, len(items))
	for _, item := range items {
		if mapped, ok := fn(item); ok {
			out = append(out, mapped)
		}
	}
	return out
}
