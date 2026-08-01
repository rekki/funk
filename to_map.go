package funk

// ToMap creates a map from a slice with a key-derivation function. If multiple
// elements produce the same key, the last element wins.
//
// For example, to index Foo values by ID:
//
//	thingForID := ToMap(xs, func(foo Foo) string { return foo.ID })
func ToMap[T any, K comparable](input []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(input))
	for _, item := range input {
		out[key(item)] = item
	}
	return out
}

// ToMapWithIndex is like ToMap, but key also receives the element's index.
func ToMapWithIndex[T any, K comparable](input []T, key func(int, T) K) map[K]T {
	out := make(map[K]T, len(input))
	for i, item := range input {
		out[key(i, item)] = item
	}
	return out
}

// ToMap2 creates a map from a slice using a function that derives both the key
// and value. If multiple elements produce the same key, the last value wins.
func ToMap2[T any, K comparable, V any](input []T, get func(T) (K, V)) map[K]V {
	out := make(map[K]V, len(input))
	for _, item := range input {
		key, value := get(item)
		out[key] = value
	}
	return out
}

// ToMap2WithIndex is like ToMap2, but get also receives the element's index.
func ToMap2WithIndex[T any, K comparable, Q any](input []T, get func(int, T) (K, Q)) map[K]Q {
	out := make(map[K]Q, len(input))
	for i, item := range input {
		k, v := get(i, item)
		out[k] = v
	}
	return out
}

// ToArrayMap groups input by a derived key, preserving input order within each
// group. The key function is called exactly once per element.
func ToArrayMap[T any, K comparable](input []T, key func(T) K) map[K][]T {
	out := make(map[K][]T, len(input))
	for _, item := range input {
		itemKey := key(item)
		out[itemKey] = append(out[itemKey], item)
	}
	return out
}
