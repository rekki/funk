package funk

import (
	"reflect"
	"testing"
)

func TestIndexBy(t *testing.T) {
	input := []string{"a", "bb", "c"}
	got := IndexBy(input, func(value string) int { return len(value) })
	want := map[int]string{1: "c", 2: "bb"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("IndexBy() = %v, want %v", got, want)
	}
}

func TestAssociate(t *testing.T) {
	input := []string{"a", "bb"}
	got := Associate(input, func(value string) (int, string) {
		return len(value), value + value
	})
	want := map[int]string{1: "aa", 2: "bbbb"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Associate() = %v, want %v", got, want)
	}
}
