package funk

import (
	"reflect"
	"testing"
)

func TestDuplicateCounts(t *testing.T) {
	if got, want := DuplicateCounts([]string{"a", "b", "b", "c", "c", "c"}), map[string]int{"b": 2, "c": 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("DuplicateCounts() = %v, want %v", got, want)
	}
}

func TestDuplicateCountsUsesComparableIdentity(t *testing.T) {
	type value struct{ name string }
	one, two, three := &value{"a"}, &value{"b"}, &value{"c"}
	input := []*value{one, one, one, two, two, three}
	want := map[*value]int{one: 3, two: 2}
	if got := DuplicateCounts(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("DuplicateCounts() = %v, want %v", got, want)
	}
}
