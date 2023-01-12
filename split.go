package funk

// `split` divide a slice into multiple subslice of
// a given maximum size
func Split[A any](from []A, size int) [][]A {
	if len(from) <= size {
		return [][]A{from}
	}

	split := [][]A{}

	chunksCount := len(from) / size
	if len(from)%size > 0 {
		chunksCount++
	}

	for i := 0; i < chunksCount; i++ {
		limit := size + i*size
		if limit > len(from) {
			limit = len(from)
		}

		split = append(split, from[i*size:limit])
	}

	return split
}
