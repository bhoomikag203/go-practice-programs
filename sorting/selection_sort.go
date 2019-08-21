package main

import (
	"fmt"
)

func selection_sort(list []int) []int {
	length := len(list)
	for i := 0; i < length-1; i++ {
		min_index := i
		for j := i + 1; j < length; j++ {
			if list[j] < list[min_index] {
				min_index = j
			}

		}

		temp := list[i]
		list[i] = list[min_index]
		list[min_index] = temp
	}

	return list
}

func main() {
	list := []int{64, 25, 12, 22, 11}
	fmt.Println(selection_sort(list))
}
