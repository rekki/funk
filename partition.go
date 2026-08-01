package funk

// Partition splits items into values accepted and rejected by predicate,
// preserving input order in both results.
func Partition[T any](items []T, predicate func(T) bool) (accepted, rejected []T) {
	for _, item := range items {
		if predicate(item) {
			accepted = append(accepted, item)
		} else {
			rejected = append(rejected, item)
		}
	}
	return accepted, rejected
}
