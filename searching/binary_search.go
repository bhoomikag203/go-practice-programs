package main

import (
	"fmt"
)

func binarySearch(list []int, key int) {
	low := 0
	high := len(list) - 1
	index := -1
	for low <= high {
		mid := (low + high) / 2

		if list[mid] == key {
			index = mid
			break
		} else if key > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if index == -1 {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Key found at position ", index+1)
	}
}
func main() {
	list := []int{1, 2, 3, 4, 5, 8}
	binarySearch(list, 1)
}
