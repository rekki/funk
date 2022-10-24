package funk

// Map returns a slice of values applying the mapper function to each element
func Map[A any, From any](from []From, mapper func(From) A) []A {
	out := make([]A, len(from))

	for i, item := range from {
		v := mapper(item)
		out[i] = v
	}

	return out
}
