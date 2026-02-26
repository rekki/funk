package funk

// Pointerify returns a pointer to v. Useful when a *T is required and you have a T.
func Pointerify[T any](v T) *T {
	return &v
}
