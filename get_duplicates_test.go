package funk

import (
	"reflect"
	"testing"
)

func TestGetDuplicates(t *testing.T) {
	if got, want := GetDuplicates([]string{"a", "b", "b", "c", "c", "c"}), map[string]int{"b": 1, "c": 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("GetDuplicates() = %v, want %v", got, want)
	}
}

func TestGetDuplicatesUsesComparableIdentity(t *testing.T) {
	type value struct{ name string }
	one, two, three := &value{"a"}, &value{"b"}, &value{"c"}
	input := []*value{one, one, one, two, two, three}
	want := map[*value]int{one: 2, two: 1}
	if got := GetDuplicates(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("GetDuplicates() = %v, want %v", got, want)
	}
}
