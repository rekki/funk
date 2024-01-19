package funk

// Default returns def if v is nil, otherwise returns *v
func Default[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
