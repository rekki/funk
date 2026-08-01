package funk

// GroupBy groups elements of a slice by a key calculated by the given getKey function.
//
// The getKey function is called for each element of the slice and the returned key is used to
// determine the group to which the element belongs.
func GroupBy[A any, B comparable](from []A, getKey func(A) B) map[B][]A {
	m := map[B][]A{}

	for _, item := range from {
		itemKey := getKey(item)
		m[itemKey] = append(m[itemKey], item)
	}
	return m
}

// GroupBySingleValue groups elements of a slice by a key calculated by the given getKey function.
//
// The getKey function is called for each element of the slice and the returned key is used to
// determine the group to which the element belongs.
//
// If multiple elements have the same key, only the last element with that key will be present in the result.
func GroupBySingleValue[A any, B comparable](from []A, getKey func(A) B) map[B]A {
	m := map[B]A{}

	for _, item := range from {
		itemKey := getKey(item)
		m[itemKey] = item
	}
	return m
}
