//implementation of merge sort
//time complexity is O(nlogn) both in worst and best case

package main

import (
	"fmt"
)

func merge_sort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) / 2
	left := merge_sort(list[:mid])
	right := merge_sort(list[mid:])
	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func main() {
	list := []int{10, 3, 2, 49, 30}
	fmt.Println(merge_sort(list))
}
