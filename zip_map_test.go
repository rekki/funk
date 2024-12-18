package funk

import (
	"reflect"
	"testing"
)

func TestZipMap(t *testing.T) {
	tests := []struct {
		keys     []string
		values   []int
		expected map[string]int
	}{
		{[]string{"a", "b", "c"}, []int{1, 2, 3}, map[string]int{"a": 1, "b": 2, "c": 3}},
		{
			[]string{"a", "b", "c", "d"},
			[]int{1, 2, 3},
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 0},
		},
		{[]string{"a"}, []int{1, 2, 3}, map[string]int{"a": 1}},
	}

	for _, test := range tests {
		actuals := ZipMap(test.keys, test.values)

		if !reflect.DeepEqual(actuals, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, actuals)
		}
	}
}
