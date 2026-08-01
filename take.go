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

// TakeRunes returns at most the first count Unicode code points in value. A
// non-positive count returns an empty string.
func TakeRunes(value string, count int) string {
	if count <= 0 {
		return ""
	}

	runes := 0
	for index := range value {
		if runes == count {
			return value[:index]
		}
		runes++
	}
	return value
}
