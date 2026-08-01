package funk

// Map applies mapper to each element of from. A nil input produces a non-nil
// empty slice.
func Map[A any, From any](from []From, mapper func(From) A) []A {
	out := make([]A, len(from))

	for i, item := range from {
		out[i] = mapper(item)
	}
	return out
}

// MapWithIndex is like Map, but mapper also receives the element's index.
func MapWithIndex[A any, From any](from []From, mapper func(From, int) A) []A {
	out := make([]A, len(from))

	for i, item := range from {
		out[i] = mapper(item, i)
	}
	return out
}
