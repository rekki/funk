package funk

import (
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	accepted, rejected := Partition([]int{0, 1, 2, 3, 4}, func(value int) bool {
		return value > 1
	})
	if want := []int{2, 3, 4}; !slices.Equal(accepted, want) {
		t.Fatalf("Partition() accepted = %v, want %v", accepted, want)
	}
	if want := []int{0, 1}; !slices.Equal(rejected, want) {
		t.Fatalf("Partition() rejected = %v, want %v", rejected, want)
	}
}
