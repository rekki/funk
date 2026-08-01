package funk

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
		from []int
		size int
		want [][]int
	}{
		{name: "empty", from: []int{}, size: 4, want: [][]int{}},
		{name: "smaller than size", from: []int{1, 2}, size: 4, want: [][]int{{1, 2}}},
		{name: "exact chunks", from: []int{1, 2, 3, 4}, size: 2, want: [][]int{{1, 2}, {3, 4}}},
		{name: "partial final chunk", from: []int{1, 2, 3, 4, 5}, size: 2, want: [][]int{{1, 2}, {3, 4}, {5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.from, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitMapEmpty(t *testing.T) {
	if got := SplitMap(map[int]string{}, 2); len(got) != 0 {
		t.Fatalf("SplitMap(empty) = %v, want no chunks", got)
	}
}
