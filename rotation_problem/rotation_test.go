package main

import "testing"
import "fmt"

func TestRotation(t *testing.T) {
	t.Run("No of steps taken to convert string1 into string2", func(t *testing.T){
		got := NoOfSteps("aab", "aab")
		want := 0
		fmt.Printf("Got %d want %d", got, want)
		if got != want {
			t.Errorf("Got %d but want %d", got, want)
		}
	})

	t.Run("Check the length of the two strings", func(t *testing.T){
		got := checkLength("aaxaabc", "aabcaax")
		want := true

		if got != want {
			t.Errorf("Got %v but want %v", got, want)
		}
	})	
}