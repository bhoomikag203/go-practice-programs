//implementation of bubble_sort
//worst case time complexity is O(n^2)

package main

import (
	"fmt"
)

func bubbleSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		sorted := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
				sorted = true
			}
		}
		if sorted == false {
			fmt.Println("given array is sorted")
			break
		}
	}
	return list
}

func main() {
	list := []int{10, 54, 20, 3, 5, 9}
	fmt.Println(bubbleSort(list))
}
