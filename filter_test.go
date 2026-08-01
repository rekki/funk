package funk

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	got := Filter([]int{0, 1, 2, 3, 4}, func(value int) bool { return value > 1 })
	want := []int{2, 3, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("Filter() = %v, want %v", got, want)
	}
	if got := Filter([]int(nil), func(int) bool { return true }); len(got) != 0 {
		t.Fatalf("Filter(nil) = %#v, want empty slice", got)
	}
	if got := Filter([]int{1}, func(int) bool { return false }); len(got) != 0 {
		t.Fatalf("Filter() = %#v, want empty slice", got)
	}
}

func TestFilterMap(t *testing.T) {
	got := FilterMap([]int{1, 2, 3}, func(value int) (string, bool) {
		return string(rune('a' + value - 1)), value%2 == 1
	})
	want := []string{"a", "c"}
	if !slices.Equal(got, want) {
		t.Fatalf("FilterMap() = %v, want %v", got, want)
	}
	if got := FilterMap([]int(nil), func(int) (int, bool) { return 0, false }); len(got) != 0 {
		t.Fatalf("FilterMap(nil) = %#v, want empty slice", got)
	}
}
