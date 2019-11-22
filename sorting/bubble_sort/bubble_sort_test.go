package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 2}, expected: []int{2, 3}},
		{input: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		output := bubbleSort(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}

}
