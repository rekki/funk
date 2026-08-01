package funk

import (
	"slices"
	"testing"
)

func TestUniquePreservesFirstOccurrenceOrder(t *testing.T) {
	if got := Unique([]string(nil)); len(got) != 0 {
		t.Fatalf("Unique(nil) = %#v, want empty slice", got)
	}

	got := Unique([]string{"b", "a", "b", "c", "a"})
	want := []string{"b", "a", "c"}
	if !slices.Equal(got, want) {
		t.Fatalf("Unique() = %v, want %v", got, want)
	}
}

func TestUniqueUsesComparableIdentity(t *testing.T) {
	type value struct{ name string }
	one, two := &value{"a"}, &value{"b"}
	got := Unique([]*value{one, two, two})
	want := []*value{one, two}
	if !slices.Equal(got, want) {
		t.Fatalf("Unique() = %v, want %v", got, want)
	}
}
