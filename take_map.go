package funk

// TakeMap returns a new map containing at most n entries from input. Which
// entries are selected is unspecified. A non-positive n returns an empty map.
func TakeMap[T any, K comparable](input map[K]T, n int) map[K]T {
	if n < 0 {
		n = 0
	}
	capacity := min(n, len(input))
	out := make(map[K]T, capacity)
	for key, value := range input {
		if len(out) == n {
			break
		}
		out[key] = value
	}
	return out
}
