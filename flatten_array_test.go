package funk

import "testing"

func TestFlattenArray(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty set returns []", args{nil}, "[]"},
		{"One item returns as expected", args{[]string{"A"}}, "[\"A\"]"},
		{"Three item returns as expected", args{[]string{"A", "B", "C"}}, "[\"A\", \"B\", \"C\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenArray(tt.args.items); got != tt.want {
				t.Errorf("FlattenArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
