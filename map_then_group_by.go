package funk

// MapThenGroupBy groups mapped elements by keys derived from the original
// elements, preserving input order within each group.
func MapThenGroupBy[A any, B comparable, C any](
	from []A,
	getKey func(A) B,
	mapFn func(A) C,
) map[B][]C {
	m := map[B][]C{}

	for _, item := range from {
		itemKey := getKey(item)
		m[itemKey] = append(m[itemKey], mapFn(item))
	}
	return m
}
