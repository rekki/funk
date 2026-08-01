package funk

import (
	"reflect"
	"testing"
)

func TestBox(t *testing.T) {
	one, two := 1, 2
	tests := []struct {
		name string
		got  []any
		want []any
	}{
		{name: "integers", got: Box(1, 2, 3), want: []any{1, 2, 3}},
		{name: "strings", got: Box("a", "b"), want: []any{"a", "b"}},
		{name: "pointers", got: Box(&one, &two), want: []any{&one, &two}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Fatalf("Box() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
