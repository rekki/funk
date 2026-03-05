package funk

import (
	"testing"
)

// TestDefaultNilPointer tests Default with a nil pointer.
func TestDefaultNilPointer(t *testing.T) {
	var intValue *int
	var strValue *string
	intDefault := 42
	strDefault := "default"

	if got := Deref(intValue, intDefault); got != intDefault {
		t.Errorf("Default(intValue, %d) = %d; want %d", intDefault, got, intDefault)
	}

	if got := Deref(strValue, strDefault); got != strDefault {
		t.Errorf("Default(strValue, %s) = %s; want %s", strDefault, got, strDefault)
	}
}

// TestDefaultNonNilPointer tests Default with a non-nil pointer.
func TestDefaultNonNilPointer(t *testing.T) {
	intValue := 10
	strValue := "hello"
	intDefault := 42
	strDefault := "default"

	if got := Deref(&intValue, intDefault); got != intValue {
		t.Errorf("Default(&intValue, %d) = %d; want %d", intDefault, got, intValue)
	}

	if got := Deref(&strValue, strDefault); got != strValue {
		t.Errorf("Default(&strValue, %s) = %s; want %s", strDefault, got, strValue)
	}
}

// TestNilPointerWithoutDefault tests Default with a nil pointer and no default value.
func TestNilPointerWithoutDefault(t *testing.T) {
	var intValue *int
	var strValue *string
	type Point struct {
		X, Y int
	}
	var pointValue *Point

	if got := Deref(intValue); got != 0 {
		t.Errorf("Default(intValue) = %d; want 0", got)
	}

	if got := Deref(strValue); got != "" {
		t.Errorf("Default(strValue) = %s; want empty string", got)
	}

	if got := Deref(pointValue); got != (Point{}) {
		t.Errorf("Default(pointValue) = %v; want zero value of Point", got)
	}
}
