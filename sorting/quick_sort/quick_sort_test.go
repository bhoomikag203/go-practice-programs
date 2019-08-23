package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 2, 4, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		output := quick_sort(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}

}
