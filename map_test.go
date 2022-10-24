package funk

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {

	type One struct {
		name string
	}

	type Two struct {
		o *One
	}

	twos := []*Two{
		{o: &One{name: "a"}},
		{o: &One{name: "b"}},
		{o: &One{name: "b"}},
	}

	ones := Map(twos, func(t *Two) One { return *t.o })

	strs := Map(ones, func(t One) string { return t.name })

	expected := "abb"
	got := strings.Join(strs, "")
	if got != expected {
		t.Fatalf("Expected to get %s but got %s", expected, got)
	}
}
