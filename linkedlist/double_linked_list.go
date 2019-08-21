package main

import (
	"errors"
	"fmt"
	// "os"
)

type node struct {
	prev *node
	item int
	next *node
}

type list struct {
	head   *node
	tail   *node
	length int
}

func (l *list) insertFirst(newNode *node) {
	if l.length == 0 {
		l.tail = newNode
		l.length++
		fmt.Printf("%d succesfully inserted\n\n", newNode.item)
	} else {
		l.head.prev = newNode
		l.length++
		fmt.Printf("%d succesfully inserted\n\n", newNode.item)
	}
	newNode.next = l.head
	l.head = newNode
	l.printList()
}

func (l *list) insertAtEnd(newNode *node) {
	if l.length == 0 {
		l.tail = newNode
		l.length++
		fmt.Printf("%d succesfully inserted\n\n", newNode.item)
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		// currentNode := l.head
		// for currentNode.next != nil {
		// 	currentNode = currentNode.next
		// }
		// currentNode.next = newNode
		// newNode.prev = currentNode
		l.length++
		fmt.Printf("%d succesfully inserted\n\n", newNode.item)
	}
	l.tail = newNode
	l.printList()
}

func (l *list) delete(item int) {
	if l.length == 0 {
		fmt.Println("list is empty")
	} else {
		currentNode := l.head
		// var previousNode *node
		for currentNode.item != item {
			if currentNode.next == nil {
				panic(errors.New("data not found"))
			}
			// previousNode = currentNode
			currentNode = currentNode.next

		}
		if currentNode == l.head {
			l.head = l.head.next
		} else {
			currentNode.prev.next = currentNode.next
		}

		if currentNode == l.tail {
			l.tail = currentNode.prev
		} else {
			currentNode.next.prev = currentNode.prev
		}

		// 	if currentNode == l.tail {
		// 		if currentNode.item == item {
		// 			previousNode = currentNode
		// 			previousNode.next = nil
		// 			previousNode = l.tail
		// 			os.Exit(1)
		// 		}
		// 	}
		// 	previousNode.next = currentNode.next
		// 	currentNode = currentNode.next
		// 	currentNode.prev = previousNode.next
		// 	l.length--
		// }

	}

	l.printList()
}

func (l *list) printList() {
	if l.length == 0 {
		fmt.Println("list is empty")
	} else {
		fmt.Print("list items are ---> ")
		currentNode := l.head
		for currentNode != nil {
			fmt.Print(currentNode.item, " ")
			currentNode = currentNode.next
		}
		fmt.Println("\n")
	}
}

func main() {
	list := &list{}

	a := node{
		item: 100,
	}

	list.insertFirst(&a)
	b := node{
		item: 200,
	}
	list.insertAtEnd(&b)

	c := node{
		item: 300,
	}
	list.insertAtEnd(&c)

	d := node{
		item: 80,
	}
	list.insertAtEnd(&d)
	e := node{
		item: 60,
	}
	list.insertFirst(&e)
	list.delete(300)

	list.delete(80)

	fmt.Println(list.head.next)

}
