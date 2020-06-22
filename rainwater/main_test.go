package rainwater

import "testing"

func TestCompute(t *testing.T) {
	// arrange
	argument := []int{1, 3, 2, 4, 1, 3, 1, 4, 5, 2, 2, 1, 4, 2, 2}
	want := 15
	// act
	got := compute(argument)
	// assert
	if got != want {
		t.Errorf("compute(%v) = %v, want %v", argument, got, want)
	}
}