package funk

import (
	"reflect"
	"strings"
	"testing"
)

func TestGroupMapBy(t *testing.T) {
	got := GroupMapBy(
		[]string{"one", "two", "three", "four"},
		func(value string) int { return len(value) },
		strings.ToUpper,
	)
	want := map[int][]string{3: {"ONE", "TWO"}, 4: {"FOUR"}, 5: {"THREE"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GroupMapBy() = %v, want %v", got, want)
	}
}
