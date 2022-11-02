package funk

// ToMapBespoke creates a map from a slice with a key-derivation and a value-derivation function
//
// eg. I have a slice of thing structs that have a field ID and field description, and I want a map of id -> description then I can do
//
//  thingForID := ToMap(xs, func (foo Foo) string { return foo.ID },  func (foo Foo) string { return foo.description })
func ToMapBespoke[T any, K comparable, V any](input []T, key func(T) K, value func(T) V) map[K]V {
	out := make(map[K]V, len(input))
	for _, item := range input {
		out[key(item)] = value(item)
	}
	return out
}
