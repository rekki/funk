package funk

// GetDuplicates returns a map of items that are duplicated in the input slice and the number of times they are duplicated
func GetDuplicates[A comparable](items []A) map[A]int {
	duplicates := map[A]int{}

	values := map[A]bool{}
	for _, item := range items {
		_, isDuplicate := values[item]
		if !isDuplicate {
			values[item] = true
		} else {
			if numRepeats, notFirstRepeat := duplicates[item]; notFirstRepeat {
				duplicates[item] = numRepeats + 1
			} else {
				duplicates[item] = 1
			}
		}
	}

	return duplicates
}
