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

	if got := Default(intValue, intDefault); got != intDefault {
		t.Errorf("Default(intValue, %d) = %d; want %d", intDefault, got, intDefault)
	}

	if got := Default(strValue, strDefault); got != strDefault {
		t.Errorf("Default(strValue, %s) = %s; want %s", strDefault, got, strDefault)
	}
}

// TestDefaultNonNilPointer tests Default with a non-nil pointer.
func TestDefaultNonNilPointer(t *testing.T) {
	intValue := 10
	strValue := "hello"
	intDefault := 42
	strDefault := "default"

	if got := Default(&intValue, intDefault); got != intValue {
		t.Errorf("Default(&intValue, %d) = %d; want %d", intDefault, got, intValue)
	}

	if got := Default(&strValue, strDefault); got != strValue {
		t.Errorf("Default(&strValue, %s) = %s; want %s", strDefault, got, strValue)
	}
}
