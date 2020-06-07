package linkedlist

import (
	"testing"
)

func TestFirstOfInitialList(t *testing.T) {
	// arrange
	var want1 interface{}
	want2 := IndexOutOfBoundsError{}
	// act
	got1, got2 := New().first()
	//assert
	if got1 != want1 || got2 != want2 {
		t.Errorf("first() = %v,%v,  want %v,%v", got1, got2, want1, want2)
	}
}

func TestLengthOfInitialList(t *testing.T) {
	// arrange
	want := 0
	// act
	got := New().length
	//assert
	if got != want {
		t.Errorf("length of initial list = %v, want %v", got, want)
	}
}
