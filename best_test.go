package funk

import "testing"

func TestBest(t *testing.T) {
	calls := 0
	got, ok := Best([]int{1, 3, 2}, func(current, next int) bool {
		calls++
		return next > current
	})
	if !ok || got != 3 || calls != 2 {
		t.Fatalf("Best() = (%d, %t) with %d calls, want (3, true) with 2 calls", got, ok, calls)
	}
}

func TestBestEmpty(t *testing.T) {
	got, ok := Best([]int(nil), func(current, next int) bool { return false })
	if ok || got != 0 {
		t.Fatalf("Best(nil) = (%d, %t), want (0, false)", got, ok)
	}
}

func TestBestDistinguishesNilItemFromEmptySlice(t *testing.T) {
	got, ok := Best([]*int{nil}, func(current, next *int) bool { return false })
	if !ok || got != nil {
		t.Fatalf("Best([nil]) = (%v, %t), want (nil, true)", got, ok)
	}
}
