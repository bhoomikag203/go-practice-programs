package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{3, 2, 4, 1}
	want := []int{1, 2, 3, 4}

	got := selection_sort(list)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
