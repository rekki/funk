package funk

func GroupBy[A any, B comparable](from []A, getKey func(A) B) map[B][]A {
	m := map[B][]A{}

	for _, item := range from {
		itemKey := getKey(item)
		if currentItems, present := m[itemKey]; present {
			m[itemKey] = append(currentItems, item)
		} else {
			m[itemKey] = []A{item}
		}
	}
	return m
}
