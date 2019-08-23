//implementation of selection sort
//time complexity is O(n^2)

package main

import (
	"fmt"
)

func main() {
	list := []int{64, 25, 12, 22, 11}
	fmt.Println(selection_sort(list))
}

func selection_sort(list []int) []int {
	length := len(list)
	for i := 0; i < length; i++ {
		min_index := i // assigns the first position as minimum element position
		for j := i + 1; j < length; j++ {
			if list[j] < list[min_index] { // finds the minimum element in each iteration
				min_index = j // updates the minimum(smaller) element's position
			}
		}
		//once the smaller element position is found out, then swap the smaller element
		list[min_index], list[i] = list[i], list[min_index]
	}
	return list
}
