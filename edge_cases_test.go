package funk

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"
	"unicode/utf8"
)

func TestDerefSupportsAnyTypeAndUsesFirstNonZeroDefault(t *testing.T) {
	var value *[]int
	first := []int{}
	second := []int{2}
	if got := Deref(value, first, second); !reflect.DeepEqual(got, first) {
		t.Fatalf("Deref() = %v, want %v", got, first)
	}
	if got := Deref((*int)(nil), 0, 42); got != 42 {
		t.Fatalf("Deref(nil, 0, 42) = %d, want 42", got)
	}

	want := []int{1}
	if got := Deref(&want, second); !reflect.DeepEqual(got, want) {
		t.Fatalf("Deref() = %v, want pointed-to value %v", got, want)
	}
}

func TestIdenticalSetsIgnoresOrderAndDuplicates(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		ys   []int
		want bool
	}{
		{name: "same set", xs: []int{1, 1, 2}, ys: []int{2, 1}, want: true},
		{name: "symmetric", xs: []int{1}, ys: []int{1, 1}, want: true},
		{name: "different set", xs: []int{1, 2}, ys: []int{1, 3}, want: false},
		{name: "nil and empty", xs: nil, ys: []int{}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdenticalSets(tt.xs, tt.ys); got != tt.want {
				t.Fatalf("IdenticalSets(%v, %v) = %t, want %t", tt.xs, tt.ys, got, tt.want)
			}
			if got := IdenticalSets(tt.ys, tt.xs); got != tt.want {
				t.Fatalf("IdenticalSets(%v, %v) = %t, want %t", tt.ys, tt.xs, got, tt.want)
			}
		})
	}
}

func TestUnionReturnsDistinctValuesInFirstOccurrenceOrder(t *testing.T) {
	got := Union([]int{1, 1, 2}, []int{2, 3, 3})
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Fatalf("Union() = %v, want %v", got, want)
	}
}

func TestToArrayMapCallsKeyOncePerItem(t *testing.T) {
	calls := 0
	got := ToArrayMap([]string{"a", "b"}, func(item string) int {
		calls++
		return len(item)
	})
	if calls != 2 {
		t.Fatalf("key called %d times, want 2", calls)
	}
	if want := []string{"a", "b"}; !slices.Equal(got[1], want) {
		t.Fatalf("ToArrayMap() = %v, want group %v", got, want)
	}
}

func TestBoundedHelpersHandleInvalidBounds(t *testing.T) {
	input := []int{1, 2, 3}
	if got := Drop(input, -1); !slices.Equal(got, input) {
		t.Fatalf("Drop() = %v, want %v", got, input)
	}
	if got := Take(input, -1); len(got) != 0 {
		t.Fatalf("Take() = %v, want empty", got)
	}
	for _, got := range [][]int{
		Range(input, 2, 1),
		Range(input, 0, -10),
		Range(input, len(input), len(input)),
	} {
		if len(got) != 0 {
			t.Fatalf("Range() = %v, want empty", got)
		}
	}
}

func TestSplitRejectsNonPositiveSizes(t *testing.T) {
	assertPanics(t, func() { Split([]int{1}, 0) })
	assertPanics(t, func() { SplitMap(map[int]int{1: 1}, -1) })
}

func TestSplitMapReturnsNewMaps(t *testing.T) {
	input := map[int]string{1: "one"}
	chunks := SplitMap(input, int(^uint(0)>>1))
	chunks[0][1] = "changed"
	if input[1] != "one" {
		t.Fatal("SplitMap() result aliases input")
	}
}

func TestSplitMapPreservesAllEntriesWithinSize(t *testing.T) {
	input := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	chunks := SplitMap(input, 2)
	combined := make(map[int]string, len(input))
	for _, chunk := range chunks {
		if len(chunk) == 0 || len(chunk) > 2 {
			t.Fatalf("chunk size = %d, want 1..2", len(chunk))
		}
		for key, value := range chunk {
			if _, exists := combined[key]; exists {
				t.Fatalf("key %d appeared in multiple chunks", key)
			}
			combined[key] = value
		}
	}
	if !reflect.DeepEqual(combined, input) {
		t.Fatalf("combined chunks = %v, want %v", combined, input)
	}
}

func TestFlattenArrayEscapesStrings(t *testing.T) {
	input := []string{`a"b`, `c\d`, "line\nnext", "\u0001"}
	formatted := FlattenArray(input)
	var got []string
	if err := json.Unmarshal([]byte(formatted), &got); err != nil {
		t.Fatalf("FlattenArray() produced invalid JSON %q: %v", formatted, err)
	}
	if !slices.Equal(got, input) {
		t.Fatalf("FlattenArray() round trip = %v, want %v", got, input)
	}
}

func TestTakeStringUsesUnicodeCodePoints(t *testing.T) {
	got := TakeString("é🙂a", 2)
	if got != "é🙂" {
		t.Fatalf("TakeString() = %q, want %q", got, "é🙂")
	}
	if !utf8.ValidString(got) {
		t.Fatalf("TakeString() returned invalid UTF-8: %x", got)
	}
}

func TestToggleStringInSetIsStable(t *testing.T) {
	if got, want := ToggleStringInSet([]string{"b", "a", "b"}, "c"), []string{"b", "a", "c"}; !slices.Equal(got, want) {
		t.Fatalf("ToggleStringInSet() = %v, want %v", got, want)
	}
	if got, want := ToggleStringInSet([]string{"b", "a", "b"}, "b"), []string{"a"}; !slices.Equal(got, want) {
		t.Fatalf("ToggleStringInSet() = %v, want %v", got, want)
	}
}

func TestTakeMapReturnsCopy(t *testing.T) {
	input := map[int]string{1: "one"}
	got := TakeMap(input, 2)
	got[1] = "changed"
	if input[1] != "one" {
		t.Fatal("TakeMap() result aliases input")
	}
	if got := TakeMap(input, -1); len(got) != 0 {
		t.Fatalf("TakeMap() = %v, want empty", got)
	}
}

func TestOnePDoesNotCompareFirstItemWithItself(t *testing.T) {
	calls := 0
	got := OneP([]int{1, 3, 2}, func(current, next int) bool {
		calls++
		return next > current
	})
	if got != 3 || calls != 2 {
		t.Fatalf("OneP() = %d with %d calls, want 3 with 2 calls", got, calls)
	}
}

func TestIntersectKeyMapAcceptsNonComparableValues(t *testing.T) {
	got := IntersectKeyMap(map[int][]string{1: {"one"}}, map[int][]string{1: {"ignored"}})
	if want := []string{"one"}; !slices.Equal(got[1], want) {
		t.Fatalf("IntersectKeyMap() = %v, want %v", got, want)
	}
}

func TestReplaceAllRecursivelyWithEmptySearch(t *testing.T) {
	if got := ReplaceAllRecursively("abc", "", "x"); got != "abc" {
		t.Fatalf("ReplaceAllRecursively() = %q, want abc", got)
	}
}

func assertPanics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("function did not panic")
		}
	}()
	fn()
}
