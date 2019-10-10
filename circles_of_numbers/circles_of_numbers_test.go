package main

import "testing"

func TestCircleOfNumbers(t *testing.T) {
	got := CircleOfNumbers(8, 6)
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
