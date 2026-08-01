package funk

// Drop returns a subslice without the first n elements. The returned slice
// shares storage with slice. If n is non-positive, Drop returns slice. If n is
// at least the length of slice, it returns an empty slice.
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}
