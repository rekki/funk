package funk

import (
	"strings"
	"testing"
)

func TestSome(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if !Some(arr, func(s string) bool { return strings.Contains(s, "a") }) {
		t.Error("Expected true")
	}
	if Some(arr, func(s string) bool { return strings.Contains(s, "d") }) {
		t.Error("Expected false")
	}

	type Point struct {
		x, y int
	}

	arr2 := []Point{{1, 2}, {3, 4}, {5, 6}, {0, 0}}
	if !Some(arr2, func(p Point) bool { return p.x == 0 }) {
		t.Error("Expected true")
	}
}
