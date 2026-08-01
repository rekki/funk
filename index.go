package funk

// IndexBy indexes items by a derived key. If multiple items produce the same
// key, the last item wins.
func IndexBy[T any, K comparable](items []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(items))
	for _, item := range items {
		out[key(item)] = item
	}
	return out
}

// Associate creates a map from keys and values derived from items. If multiple
// items produce the same key, the last value wins.
func Associate[T any, K comparable, V any](items []T, get func(T) (K, V)) map[K]V {
	out := make(map[K]V, len(items))
	for _, item := range items {
		key, value := get(item)
		out[key] = value
	}
	return out
}
