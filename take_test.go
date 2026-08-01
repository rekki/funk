package funk

import (
	"reflect"
	"testing"
	"unicode/utf8"
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
		if got := Take(test.items, test.count); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Take(%v, %d) = %v, want %v", test.items, test.count, got, test.want)
		}
	}
}

func TestTakeRunes(t *testing.T) {
	got := TakeRunes("é🙂a", 2)
	if got != "é🙂" {
		t.Fatalf("TakeRunes() = %q, want %q", got, "é🙂")
	}
	if !utf8.ValidString(got) {
		t.Fatalf("TakeRunes() returned invalid UTF-8: %x", got)
	}
	if got := TakeRunes("abc", 0); got != "" {
		t.Fatalf("TakeRunes() = %q, want empty", got)
	}
}
