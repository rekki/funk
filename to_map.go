package funk

// ToMap creates a map from a slice with a key-derivation function
//
// eg. I have a slice of thing structs that have a field ID and I want a map of id -> thing. then I can do
//
//	thingForID := ToMap(xs, func (foo Foo) string { return foo.ID })
func ToMap[T any, K comparable](input []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(input))
	for _, item := range input {
		out[key(item)] = item
	}
	return out
}

// ToArrayMap creates a map from a slice with a key-derivation function
//Useful where the key may be duplicated

func ToArrayMap[T any, K comparable](input []T, key func(T) K) map[K][]T {
	out := make(map[K][]T, len(input))
	for _, item := range input {
		if values, exist := out[key(item)]; exist {
			out[key(item)] = append(values, item)
		} else {
			out[key(item)] = []T{item}
		}
	}
	return out
}
