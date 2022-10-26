package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestOneP(t *testing.T) {

	numbers := []numberClass{
		{2, "prime"},
		{3, "prime"},
		{5, "prime"}}

	expected := numberClass{5, "prime"}
	selectedOne := OneP(numbers, func(i numberClass, j numberClass) bool {
		return j.number > i.number
	})

	if !reflect.DeepEqual(selectedOne, expected) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(selectedOne), pretty.Sprint(expected))
	}
}
