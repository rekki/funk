package funk

// MapThenGroupBy groups elements of a slice by a key calculated by the given getKey function.
func MapThenGroupBy[A any, B comparable, C any](
	from []A,
	getKey func(A) B,
	mapFn func(A) C,
) map[B][]C {
	m := map[B][]C{}

	for _, item := range from {
		itemKey := getKey(item)
		if currentItems, present := m[itemKey]; present {
			m[itemKey] = append(currentItems, mapFn(item))
		} else {
			m[itemKey] = []C{mapFn(item)}
		}
	}
	return m
}
