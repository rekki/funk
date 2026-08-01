package funk

// ToggleStringInSet returns the distinct strings in input with target toggled.
// It preserves first-occurrence order; a newly added target is appended.
func ToggleStringInSet(input []string, target string) []string {
	result := make([]string, 0, len(input)+1)
	seen := make(map[string]struct{}, len(input))
	foundTarget := false
	for _, item := range input {
		if item == target {
			foundTarget = true
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	if !foundTarget {
		result = append(result, target)
	}
	return result
}
