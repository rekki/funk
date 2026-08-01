package funk

// Split divides from into subslices containing at most size elements. The
// subslices share storage with from. Empty input produces no chunks. Split
// panics if size is not positive.
func Split[A any](from []A, size int) [][]A {
	if size <= 0 {
		panic("funk: Split size must be positive")
	}
	if len(from) == 0 {
		return [][]A{}
	}
	if len(from) <= size {
		return [][]A{from}
	}

	chunks := make([][]A, 0, (len(from)-1)/size+1)
	for start := 0; start < len(from); start += size {
		end := min(start+size, len(from))
		chunks = append(chunks, from[start:end])
	}
	return chunks
}

// SplitMap divides from into new maps containing at most size entries. Map
// iteration order, and therefore chunk membership, is unspecified. Empty input
// produces no chunks. SplitMap panics if size is not positive.
func SplitMap[K comparable, V any](from map[K]V, size int) []map[K]V {
	if size <= 0 {
		panic("funk: SplitMap size must be positive")
	}
	if len(from) == 0 {
		return []map[K]V{}
	}
	chunkCapacity := min(size, len(from))
	splitMaps := make([]map[K]V, 0, (len(from)-1)/size+1)
	subMap := make(map[K]V, chunkCapacity)
	for k, v := range from {
		if len(subMap) == size {
			splitMaps = append(splitMaps, subMap)
			subMap = make(map[K]V, chunkCapacity)
		}
		subMap[k] = v
	}
	return append(splitMaps, subMap)
}
