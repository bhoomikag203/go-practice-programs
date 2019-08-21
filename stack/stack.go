package main

import (
	"fmt"
	"os"
)

var top = -1

const Max = 2

var stack []int

func push(ele int) {
	fmt.Println(top)
	if top == Max-1 {
		fmt.Print("cannot push, stack is full")
		os.Exit(1)
	}
	top++
	stack = append(stack, ele)
	fmt.Print(stack)
}

func pop() {
	fmt.Println(top)
	if top == -1 {
		fmt.Println("Cannot pop, stack is empty")
		os.Exit(1)
	}
	top--
	stack = stack[:len(stack)-1]
}

func main() {
	push(5)
	push(10)
	pop()
	push(4)
	push(40)
}
