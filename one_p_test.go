package funk

import (
	"reflect"
	"testing"
)

func TestOnePPanicsOnEmptySlice(t *testing.T) {
	assertPanics(t, func() {
		OneP([]int{}, func(current, next int) bool { return false })
	})
}

func TestOneP(t *testing.T) {
	numbers := []numberClass{
		{2, "prime"},
		{3, "prime"},
		{5, "prime"},
	}

	expected := numberClass{5, "prime"}
	selectedOne := OneP(numbers, func(i numberClass, j numberClass) bool {
		return j.number > i.number
	})

	if !reflect.DeepEqual(selectedOne, expected) {
		t.Fatalf("OneP() = %v, want %v", selectedOne, expected)
	}
}

func TestOneRef(t *testing.T) {
	one := 1
	two := 2
	three := 3
	numbers := []*int{&one, &two, &three}
	expected := &one

	found := OneRef(numbers)

	if !reflect.DeepEqual(found, expected) {
		t.Fatalf("OneRef() = %v, want %v", found, expected)
	}

	emptyNumbers := []*int{}
	var expectedEmpty *int

	foundEmpty := OneRef(emptyNumbers)

	if !reflect.DeepEqual(foundEmpty, expectedEmpty) {
		t.Fatalf("OneRef() = %v, want %v", foundEmpty, expectedEmpty)
	}
}
