//implementation of quick sort
//time complexity is O(n^2)
//best casse time complexity is O(nlogn)

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
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func main() {
	list := []int{10, 3, 2, 49, 30}
	fmt.Println(quickSort(list, 0, len(list)-1))
}
