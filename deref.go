package funk

import "reflect"

// Deref returns the value pointed to by v. If v is nil, it returns the first
// non-zero default value, or the zero value of T when none is provided.
func Deref[T any](v *T, defaults ...T) T {
	if v != nil {
		return *v
	}
	for _, value := range defaults {
		if !reflect.ValueOf(&value).Elem().IsZero() {
			return value
		}
	}
	var zero T
	return zero
}
