package funk

// GetDuplicates returns values that occur more than once. Each result count is
// the number of occurrences after the first.
func GetDuplicates[A comparable](items []A) map[A]int {
	counts := make(map[A]int, len(items))
	for _, item := range items {
		counts[item]++
	}

	duplicates := make(map[A]int)
	for item, count := range counts {
		if count > 1 {
			duplicates[item] = count - 1
		}
	}
	return duplicates
}
