package main

import (
	"fmt"
)

func quickSort(list []int, low, high int) []int {
	if low > high {
		return list
	}
	p := partition(list, low, high)
	quickSort(list, low, p-1)
	quickSort(list, p+1, high)
	return list
}

func partition(list []int, low, high int) int {
	pivot := list[high]
	for j := low; j < high; j++ {
		if list[j] < pivot {
			list[j], list[low] = list[low], list[j]
			low++
		}
	}
	list[low], list[high] = list[high], list[low]
	return low
}

func main() {
	list := []int{10, 3, 2, 49, 30}
	fmt.Println(quickSort(list, 0, len(list)-1))
}
