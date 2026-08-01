package funk

import (
	"reflect"
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	got := Difference([]string{"A", "A", "B", "C"}, []string{"C", "D"})
	want := []string{"A", "A", "B"}
	if !slices.Equal(got, want) {
		t.Fatalf("Difference() = %v, want %v", got, want)
	}
	if got := Difference([]string{"A"}, []string{"A"}); got != nil {
		t.Fatalf("Difference() = %#v, want nil", got)
	}
}

func TestIntersect(t *testing.T) {
	got := Intersect([]string{"A", "B", "B", "C"}, []string{"B", "D"})
	want := []string{"B", "B"}
	if !slices.Equal(got, want) {
		t.Fatalf("Intersect() = %v, want %v", got, want)
	}
	if got := Intersect([]string{"A"}, nil); got != nil {
		t.Fatalf("Intersect() = %#v, want nil", got)
	}
}

func TestSetEqual(t *testing.T) {
	tests := []struct {
		xs, ys []int
		want   bool
	}{
		{xs: []int{1, 1, 2}, ys: []int{2, 1}, want: true},
		{xs: []int{1}, ys: []int{1, 1}, want: true},
		{xs: []int{1, 2}, ys: []int{1, 3}, want: false},
		{xs: nil, ys: []int{}, want: true},
	}
	for _, test := range tests {
		if got := SetEqual(test.xs, test.ys); got != test.want {
			t.Errorf("SetEqual(%v, %v) = %t, want %t", test.xs, test.ys, got, test.want)
		}
		if got := SetEqual(test.ys, test.xs); got != test.want {
			t.Errorf("SetEqual(%v, %v) = %t, want %t", test.ys, test.xs, got, test.want)
		}
	}

	if !reflect.DeepEqual(Difference([]int{1}, nil), []int{1}) {
		t.Fatal("Difference changed the non-empty nil-right behavior")
	}
}
