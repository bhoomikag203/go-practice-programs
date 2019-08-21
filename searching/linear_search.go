package main

import (
	"fmt"
)

func linearSearch(list []int, key int) {
	index := -1
	for i := 0; i < len(list); i++ {
		if list[i] == key {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Key found at position ", index+1)
	}
}

func main() {
	list := []int{10, 25, 46}
	linearSearch(list, 7)
}
