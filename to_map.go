package funk

/**
 * ToMap converts an array to a map
 */
func ToMap[T any, K comparable](input []T, key func(T) K) map[K]T {
	out := map[K]T{}
	for _, item := range input {
		out[key(item)] = item
	}
	return out
}
