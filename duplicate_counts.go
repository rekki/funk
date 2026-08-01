package funk

// DuplicateCounts returns the total occurrence count for each value that
// appears more than once.
func DuplicateCounts[T comparable](items []T) map[T]int {
	counts := make(map[T]int, len(items))
	for _, item := range items {
		counts[item]++
	}

	for item, count := range counts {
		if count == 1 {
			delete(counts, item)
			continue
		}
		counts[item] = count
	}
	return counts
}
