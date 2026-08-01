package funk

// PartitionBy splits from into values accepted and rejected by predicate,
// preserving input order in both results.
func PartitionBy[A any](from []A, filter func(A) bool) (left []A, right []A) {
	var l []A
	var r []A

	for _, item := range from {
		if filter(item) {
			l = append(l, item)
		} else {
			r = append(r, item)
		}
	}

	return l, r
}
