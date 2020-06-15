package distributecandies

import (
	"testing"
)

func TestDistributeCandies1(t *testing.T) {
	// arrange
	argument := []int{1, 1, 2, 2, 3, 3}
	want := 3
	// act
	got := distributeCandies(argument)
	// assert
	if got != want {
		t.Errorf("distributeCandies(%v) = %v, want %v", argument, got, want)
	}
}

func TestDistributeCandies2(t *testing.T) {
	// arrange
	argument := []int{1, 1, 2, 3}
	want := 2
	// act
	got := distributeCandies(argument)
	// assert
	if got != want {
		t.Errorf("distributeCandies(%v) = %v, want %v", argument, got, want)
	}
}
