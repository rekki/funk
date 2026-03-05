package funk

import "cmp"

// Deref returns def if v is nil, otherwise returns *v
func Deref[T comparable](v *T, def ...T) T {
	if v == nil {
		return cmp.Or(def...)
	}
	return *v
}
