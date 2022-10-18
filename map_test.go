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

func TestUniqueStrings(t *testing.T) {

	letters := []string{"a", "b", "b"}

	dedup := Unique(letters)

	expected := "ab"
	got := strings.Join(dedup, "")
	if got != expected {
		t.Fatalf("Expected to get %s but got %s", expected, got)
	}
}

func TestUniqueStructs(t *testing.T) {
	type abc struct {
		n string
	}

	one := abc{"a"}
	two := abc{"b"}

	letters := []*abc{
		&one,
		&two,
		&two,
	}

	expected := []*abc{
		&one,
		&two,
	}

	dedup := Unique(letters)

	if len(dedup) != len(expected) {
		t.Fatalf("Expected to get %d elements but got %d", len(expected), len(dedup))
	}

	for idx := range dedup {
		if dedup[idx] != expected[idx] {
			t.Errorf("Element idx = %d doesn't match", idx)
		}
	}
}
