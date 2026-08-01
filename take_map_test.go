package funk

import "testing"

func TestTakeMapSizeAndMembership(t *testing.T) {
	input := map[int]string{1: "one", 2: "two", 3: "three"}
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "negative", n: -1, want: 0},
		{name: "zero", n: 0, want: 0},
		{name: "subset", n: 2, want: 2},
		{name: "equal", n: 3, want: 3},
		{name: "greater", n: 4, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeMap(input, tt.n)
			if len(got) != tt.want {
				t.Fatalf("len(TakeMap()) = %d, want %d", len(got), tt.want)
			}
			for key, value := range got {
				if input[key] != value {
					t.Fatalf("TakeMap() contains unexpected entry %d: %q", key, value)
				}
			}
		})
	}
}
