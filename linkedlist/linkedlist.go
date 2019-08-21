package main

import (
	"fmt"
)

type numbers struct {
	number int
	next   *numbers
}

func addNodeEnd(newNode, numberList *numbers) *numbers {
	if numberList == nil {
		return numberList
	}

	for n := numberList; n != nil; n = n.next {
		if n.next == nil {
			n.next = newNode
			return numberList
		}
	}
	return numberList
}

// func addNodeBegining(newNode, numberList *numbers) *numbers {
// 	if numberList != nil {
// 		numberList = newNode
// 	}
// 	newNode.next = numberList
// 	return numberList

// }

func deleteNode(number int, numberList *numbers) *numbers {
	if numberList == nil {
		return nil
	}

	n := numberList

	if n.number == number {
		n = n.next
		return n
	}

	for n != nil {
		if n.next.number == number {
			n.next = n.next.next
			return numberList
		}
		n = n.next
	}

	return numberList
}

func reverse(numberList *numbers) *numbers {
	if numberList == nil {
		return numberList
	}

	n := numberList

	if n.next == nil {
		return n
	} else {
		head := reverse(n.next)
		n.next.next = n
		n.next = nil
		return head
	}
}

func printList(numberList *numbers) {
	for n := numberList; n != nil; n = n.next {
		fmt.Println(n)
	}
	fmt.Println()
}

func main() {
	number := &numbers{10, nil}
	numberList := number
	fmt.Println(number)

	number1 := &numbers{20, nil}
	fmt.Println(number1)

	number2 := &numbers{30, nil}
	fmt.Println(number2)

	number3 := &numbers{40, nil}
	fmt.Println(number3)
	fmt.Println()

	numberList = addNodeEnd(number1, numberList)
	printList(numberList)

	numberList = addNodeEnd(number2, numberList)
	printList(numberList)

	numberList = reverse(numberList)
	printList(numberList)

	numberList = deleteNode(20, numberList)
	printList(numberList)

	numberList = addNodeEnd(number1, numberList)
	printList(numberList)

	numberList = addNodeEnd(number2, numberList)
	printList(numberList)
	// numberList = addNodeBegining(number3, numberList)
	// printList(numberList)

}
