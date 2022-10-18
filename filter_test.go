package funk

import (
	"reflect"
	"testing"

	"github.com/rekki/go/pkg/util"
)

func TestFilter(t *testing.T) {

	ints := []int{0, 1, 2, 3, 4}

	expected := []int{2, 3, 4}
	filtered := Filter(ints, func(i int) bool {
		return i > 1
	})

	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("Expected to get %s but got %s", util.PPs(filtered), util.PPs(expected))
	}
}

type numberClass struct {
	number int
	label  string
}
