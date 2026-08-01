package funk

import "testing"

func TestReduce(t *testing.T) {
	if got := Reduce([]int{1, 2, 3}, func(total, value int) int { return total + value }, 10); got != 16 {
		t.Fatalf("Reduce() = %d, want 16", got)
	}
}
