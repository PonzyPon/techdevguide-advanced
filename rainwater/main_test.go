package rainwater

import "testing"

func TestTrap(t *testing.T) {
	// arrange
	argument := []int{1, 3, 2, 4, 1, 3, 1, 4, 5, 2, 2, 1, 4, 2, 2}
	want := 15
	// act
	got := trap(argument)
	// assert
	if got != want {
		t.Errorf("trap(%v) = %v, want %v", argument, got, want)
	}
}
