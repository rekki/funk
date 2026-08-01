package funk

import (
	"reflect"
	"testing"
)

func TestFilterToEmptyIncludesMatches(t *testing.T) {
	got := FilterToEmpty([]int{1, 2}, func(value int) bool { return value == 2 })
	want := []int{2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FilterToEmpty() = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}

	expected := []int{2, 3, 4}
	filtered := Filter(ints, func(i int) bool {
		return i > 1
	})

	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("Filter() = %v, want %v", filtered, expected)
	}
}

type numberClass struct {
	number int
	label  string
}
