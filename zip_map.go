package funk

// ZipMap pairs keys with values at matching indexes. Missing values use the
// zero value of V, surplus values are ignored, and the last duplicate key wins.
func ZipMap[K comparable, V any](keys []K, values []V) map[K]V {
	m := make(map[K]V)

	for i, k := range keys {
		var value V
		if i < len(values) {
			value = values[i]
		}
		m[k] = value
	}

	return m
}
