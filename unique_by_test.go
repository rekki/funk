package funk

import (
	"slices"
	"testing"
)

func TestUniqueByPreservesFirstOccurrenceOrder(t *testing.T) {
	if got := UniqueBy([]string(nil), func(value string) int { return len(value) }); len(got) != 0 {
		t.Fatalf("UniqueBy(nil) = %#v, want empty slice", got)
	}

	got := UniqueBy([]string{"b", "aa", "c", "dd"}, func(value string) int { return len(value) })
	want := []string{"b", "aa"}
	if !slices.Equal(got, want) {
		t.Fatalf("UniqueBy() = %v, want %v", got, want)
	}
}

func TestUniqueByKeepsFirstItemForKey(t *testing.T) {
	type value struct {
		id   int
		name string
	}
	first := &value{id: 1, name: "first"}
	second := &value{id: 1, name: "second"}
	third := &value{id: 2, name: "third"}
	got := UniqueBy([]*value{first, second, third}, func(item *value) int { return item.id })
	want := []*value{first, third}
	if !slices.Equal(got, want) {
		t.Fatalf("UniqueBy() = %v, want %v", got, want)
	}
}
