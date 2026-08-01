package funk

// Reduce folds items from left to right, starting with initialValue.
func Reduce[T, R any](items []T, reducer func(R, T) R, initialValue R) R {
	acc := initialValue

	for _, item := range items {
		acc = reducer(acc, item)
	}

	return acc
}
