package funk

import (
	"reflect"
	"testing"
)

const SIZE = 4

func TestSplit_SliceSmallerThanLength(t *testing.T) {
	from := []string{"one", "two"}
	want := [][]string{{"one", "two"}}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice smaller than size]", got, want)
	}
}

func TestSplit_SliceSplitPerfectly(t *testing.T) {
	from := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice split perfectly]", got, want)
	}
}

func TestSplit_SliceDontSplitPerfectly(t *testing.T) {
	from := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10},
	}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice don't split perfectly]", got, want)
	}
}
