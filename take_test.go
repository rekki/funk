package funk

import (
	"slices"
	"testing"
)

func TestTake(t *testing.T) {
	tests := []struct {
		items []string
		count int
		want  []string
	}{
		{items: []string{"a", "b", "c"}, count: 2, want: []string{"a", "b"}},
		{items: []string{"a"}, count: 2, want: []string{"a"}},
		{items: []string{"a"}, count: 0, want: []string{}},
		{items: nil, count: 3, want: nil},
	}
	for _, test := range tests {
		if got := Take(test.items, test.count); !slices.Equal(got, test.want) {
			t.Errorf("Take(%v, %d) = %v, want %v", test.items, test.count, got, test.want)
		}
	}
}
