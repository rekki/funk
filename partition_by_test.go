package funk

import (
	"reflect"
	"testing"
)

func TestPartitionBy(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}

	expectedL := []int{2, 3, 4}
	expectedR := []int{0, 1}
	partitionByL, partitionByR := PartitionBy(ints, func(i int) bool {
		return i > 1
	})

	if !reflect.DeepEqual(expectedL, partitionByL) {
		t.Fatalf("PartitionBy() left = %v, want %v", partitionByL, expectedL)
	}
	if !reflect.DeepEqual(expectedR, partitionByR) {
		t.Fatalf("PartitionBy() right = %v, want %v", partitionByR, expectedR)
	}
}
