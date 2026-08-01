package funk

// GetDuplicates returns values that occur more than once. Each result count is
// the number of occurrences after the first.
func GetDuplicates[T comparable](items []T) map[T]int {
	counts := make(map[T]int, len(items))
	for _, item := range items {
		counts[item]++
	}

	for item, count := range counts {
		if count == 1 {
			delete(counts, item)
			continue
		}
		counts[item] = count - 1
	}
	return counts
}
