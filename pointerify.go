package funk

// Pointerify returns a pointer to the given value.
func Pointerify[A any](a A) *A {
	return &a
}
