package main

import (
	"errors"
	"fmt"
)

func insertEle(pos, ele int, numbers []int) []int {

	var length int
	length = len(numbers)
	if pos > length {
		panic(errors.New("cannot insert an element as position is greater than array length"))
	}
	numbers = append(numbers, 0)
	copy(numbers[pos+1:], numbers[pos:])
	numbers[pos] = ele
	fmt.Printf("%d is inserted", ele)
	fmt.Println()
	return numbers
}

func deleteEle(pos int, numbers []int) []int {
	var length int
	length = len(numbers)
	if pos > length {
		panic(errors.New("cannot insert an element as position is greater than array length"))
	}
	ele := numbers[pos]
	copy(numbers[pos:], numbers[pos+1:])
	fmt.Printf("%d is deleted \n", ele)
	return numbers[:len(numbers)-1]
}

func search(key int, numbers []int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == key {
			fmt.Printf("%d found at %d", key, i)
			fmt.Println()
			break
		} else {
			fmt.Printf("%d not found ", key)
			fmt.Println()
			break
		}
	}
}

func sort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		flag := 0
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
				flag = 1
			}
		}
		if flag == 0 {
			break
		}

	}
	return numbers
}

func main() {
	numbers := []int{100}
	fmt.Println(numbers)
	numbers = insertEle(1, 30, numbers)
	fmt.Println(numbers)
	numbers = insertEle(2, 50, numbers)
	fmt.Println(numbers)
	numbers = insertEle(2, 10, numbers)
	fmt.Println(numbers)
	numbers = deleteEle(1, numbers)
	fmt.Println(numbers)
	numbers = insertEle(1, 30, numbers)
	fmt.Println(numbers)
	search(30, numbers)
	fmt.Println(numbers)
	numbers = insertEle(2, 10, numbers)
	numbers = sort(numbers)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
