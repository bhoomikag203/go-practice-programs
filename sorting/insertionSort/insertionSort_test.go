package insertion_sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 2}, expected: []int{2, 3}},
		{input: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		output := insertion_sort(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}
}
