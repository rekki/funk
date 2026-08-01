package funk

import "testing"

func TestAll(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  bool
	}{
		{name: "empty", items: nil, want: true},
		{name: "all accepted", items: []string{"a", "b"}, want: true},
		{name: "one rejected", items: []string{"a", ""}, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := All(test.items, func(value string) bool { return value != "" }); got != test.want {
				t.Fatalf("All() = %t, want %t", got, test.want)
			}
		})
	}
}
