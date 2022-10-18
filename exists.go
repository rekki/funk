package funk

//Exists is used to check whether at least one item in a collection satisfies the predicate
func Exists[A any](from []A, condition func(A) bool) bool {
	for _, item := range from {
		if condition(item) {
			return true
		}
	}
	return false
}
