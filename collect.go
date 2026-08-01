package funk

// Collect returns mapped values for elements accepted by filter. It is the
// equivalent of Filter followed by Map, but traverses the input once.
func Collect[A any, From any](from []From, filter func(From) bool, mapper func(From) A) []A {
	out := make([]A, 0, len(from))

	for _, item := range from {
		if filter(item) {
			out = append(out, mapper(item))
		}
	}
	return out
}

// Collect2 combines filtering and mapping in one callback. The callback's
// boolean reports whether its mapped value should be included.
func Collect2[A any, From any](from []From, filter func(From) (bool, A)) []A {
	out := make([]A, 0, len(from))

	for _, item := range from {
		if passed, mappedValue := filter(item); passed {
			out = append(out, mappedValue)
		}
	}
	return out
}
