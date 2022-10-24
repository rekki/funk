package funk

// ToMap creates a map from a slice with a key-derivation function
//
// eg. I have a slice of thing structs that have a field ID and I want a map of id -> thing. then I can do
//
//  thingForID := ToMap(xs, func (foo Foo) string { return foo.ID })
func ToMap[T any, K comparable](input []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(input))
	for _, item := range input {
		out[key(item)] = item
	}
	return out
}
