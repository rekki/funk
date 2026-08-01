package funk

func ToggleStringInSet(strings []string, target string) []string {
	// Create a map to deduplicate the strings
	stringSet := make(map[string]bool)
	for _, str := range strings {
		stringSet[str] = true
	}

	// Flip target string in the set
	if stringSet[target] {
		delete(stringSet, target) // Remove if it exists
	} else {
		stringSet[target] = true // Add if it does not exist
	}

	// Convert map back to a slice
	result := make([]string, 0, len(stringSet))
	for str := range stringSet {
		result = append(result, str)
	}

	return result
}
