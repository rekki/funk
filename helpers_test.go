package funk

import (
	"reflect"
	"slices"
	"testing"
)

func TestPreviouslyUncoveredHelpers(t *testing.T) {
	if got, want := Collect2([]int{1, 2, 3}, func(value int) (bool, string) {
		return value%2 == 1, string(rune('a' + value - 1))
	}), []string{"a", "c"}; !slices.Equal(got, want) {
		t.Fatalf("Collect2() = %v, want %v", got, want)
	}

	if got := FilterToEmpty([]int{1, 2}, func(int) bool { return false }); got == nil || len(got) != 0 {
		t.Fatalf("FilterToEmpty() = %#v, want non-nil empty slice", got)
	}

	if got, want := Remove([]int{1, 2, 3}, []int{2}), []int{1, 3}; !slices.Equal(got, want) {
		t.Fatalf("Remove() = %v, want %v", got, want)
	}

	if got, want := GroupBySingleValue([]string{"a", "bb", "c"}, func(value string) int { return len(value) }), map[int]string{1: "c", 2: "bb"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("GroupBySingleValue() = %v, want %v", got, want)
	}

	if got, want := MapWithIndex([]string{"a", "b"}, func(value string, index int) string {
		return value + string(rune('0'+index))
	}), []string{"a0", "b1"}; !slices.Equal(got, want) {
		t.Fatalf("MapWithIndex() = %v, want %v", got, want)
	}

	one, two := 1, 2
	if got := OneOrNil([]*int{&one, &two}, func(current, next *int) bool { return *next > *current }); got != &two {
		t.Fatalf("OneOrNil() = %v, want pointer to 2", got)
	}
	if got := OneOrNil[int](nil, func(current, next *int) bool { return false }); got != nil {
		t.Fatalf("OneOrNil(nil) = %v, want nil", got)
	}

	printed := PrintMap(map[string]int{"one": 1, "two": 2})
	slices.Sort(printed)
	if want := []string{"key: one; value: 1", "key: two; value: 2"}; !slices.Equal(printed, want) {
		t.Fatalf("PrintMap() = %v, want %v", printed, want)
	}

	if got := Reduce([]int{1, 2, 3}, func(total, value int) int { return total + value }, 10); got != 16 {
		t.Fatalf("Reduce() = %d, want 16", got)
	}
}

func TestMapConversionVariants(t *testing.T) {
	input := []string{"a", "bb"}
	if got, want := ToMapWithIndex(input, func(index int, value string) int { return index + len(value) }), map[int]string{1: "a", 3: "bb"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ToMapWithIndex() = %v, want %v", got, want)
	}
	if got, want := ToMap2(input, func(value string) (int, string) { return len(value), value + value }), map[int]string{1: "aa", 2: "bbbb"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ToMap2() = %v, want %v", got, want)
	}
	if got, want := ToMap2WithIndex(input, func(index int, value string) (int, int) { return index, len(value) }), map[int]int{0: 1, 1: 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ToMap2WithIndex() = %v, want %v", got, want)
	}
}
