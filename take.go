package funk

// Take returns a subslice containing at most the first count elements of items.
// A non-positive count returns an empty subslice. The result shares storage
// with items.
func Take[T any](items []T, count int) []T {
	if count <= 0 {
		return items[:0]
	}
	if count >= len(items) {
		return items
	}
	return items[:count]
}
