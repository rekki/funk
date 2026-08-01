package funk

import (
	"slices"
	"testing"
)

func TestSetZeroValueAndMutations(t *testing.T) {
	var set Set[int]
	set.Add(2)
	set.AddMulti(1, 2, 3)
	if !set.Has(2) || set.Size() != 3 {
		t.Fatalf("set = %v (size %d), want values 1, 2, 3", set.Values(), set.Size())
	}
	set.Remove(2)
	if set.Has(2) || set.Size() != 2 {
		t.Fatalf("set after Remove = %v", set.Values())
	}
	set.Clear()
	if set.Size() != 0 {
		t.Fatalf("set size after Clear = %d, want 0", set.Size())
	}
}

func TestSetOperations(t *testing.T) {
	left := NewSet(1, 2, 3)
	right := NewSet(3, 4)

	assertSetValues(t, left.Filter(func(v int) bool { return v%2 == 1 }), []int{1, 3})
	assertSetValues(t, left.Union(right), []int{1, 2, 3, 4})
	assertSetValues(t, left.Intersect(right), []int{3})
	assertSetValues(t, left.Difference(right), []int{1, 2})

	if left.Size() != 3 || right.Size() != 2 {
		t.Fatal("set operation mutated an input")
	}
}

func assertSetValues[T comparable](t *testing.T, set *Set[T], want []T) {
	t.Helper()
	got := set.Values()
	if len(got) != len(want) {
		t.Fatalf("Values() = %v, want %v", got, want)
	}
	for _, value := range want {
		if !slices.Contains(got, value) {
			t.Fatalf("Values() = %v, missing %v", got, value)
		}
	}
}
