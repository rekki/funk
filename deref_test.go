package funk

import "testing"

func TestDerefNilPointerUsesDefault(t *testing.T) {
	var intValue *int
	var strValue *string
	if got := Deref(intValue, 42); got != 42 {
		t.Fatalf("Deref(nil, 42) = %d, want 42", got)
	}
	if got := Deref(strValue, "default"); got != "default" {
		t.Fatalf("Deref(nil, default) = %q, want default", got)
	}
}

func TestDerefNonNilPointerUsesPointedToValue(t *testing.T) {
	value := 10
	if got := Deref(&value, 42); got != value {
		t.Fatalf("Deref(&value, 42) = %d, want %d", got, value)
	}
}

func TestDerefNilPointerWithoutDefaultUsesZeroValue(t *testing.T) {
	type point struct{ x, y int }
	var value *point
	if got := Deref(value); got != (point{}) {
		t.Fatalf("Deref(nil) = %v, want zero value", got)
	}
}
