package funk

import (
	"slices"
	"testing"
)

func TestToggleStringInSet(t *testing.T) {
	type args struct {
		strings []string
		target  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test with empty slice", args{strings: []string{}, target: "test"}, []string{"test"}},
		{"Test with existing target", args{strings: []string{"test"}, target: "test"}, []string{}},
		{
			"Test with non-existing target",
			args{strings: []string{"test"}, target: "new"},
			[]string{"test", "new"},
		},
		{
			"Test with dupes",
			args{strings: []string{"test", "test"}, target: "new"},
			[]string{"test", "new"},
		},
		{
			"Test with duplicated target",
			args{strings: []string{"test", "test", "new"}, target: "test"},
			[]string{"new"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleStringInSet(
				tt.args.strings,
				tt.args.target,
			); !slices.Equal(got, tt.want) {
				t.Errorf("ToggleStringInSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
