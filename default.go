package funk

import "reflect"

// Default returns def if v is nil, otherwise returns *v
func Default[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

// ZeroValueDefault returns def if v is a zero value, otherwise returns v
func ZeroValueDefault[T any](v T, def T) T {
	val := reflect.ValueOf(v)
	if val.IsZero() {
		return def
	}
	return v
}
