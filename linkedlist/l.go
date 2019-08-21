package main

import (
	"errors"
	"fmt"
	"os"
)

type node struct {
	data int
	next *node
}

type list struct {
	head   *node
	length int
}

func (l *list) append(newNode *node) {
	if l.length == 0 {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.length++
	fmt.Println(newNode.data, "succesfully inserted")
	l.printList()
}

func (l *list) delete(data int) {
	if l.length == 0 {
		panic(errors.New("list is empty"))
	}
	//to delete the first item
	if l.head.data == data {
		fmt.Printf("%d is deleted\n", data)
		l.head = l.head.next
		l.length--
		l.printList()

	} else {
		//to delete otherthan first element
		var previousNode *node
		currentNode := l.head
		for currentNode.data != data {
			if currentNode.next == nil {
				panic(errors.New("data not found"))
			}
			previousNode = currentNode
			currentNode = currentNode.next
		}
		fmt.Printf("%d is deleted\n", data)
		previousNode.next = currentNode.next
		l.length--
		l.printList()
	}

}

func (l *list) printList() {
	currentNode := l.head
	fmt.Print("list elements --->")

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *list) insertFirst(newNode *node) {
	if l.length == 0 {
		l.head = newNode
	} else {
		var previousNode *node
		currentNode := l.head
		previousNode = newNode
		previousNode.next = currentNode
		l.head = previousNode
		l.length++
	}
	l.printList()
}

func (l *list) reverseList() {
	if l.length == 0 {
		panic(errors.New("list is empty"))
	} else {
		currentNode := l.head
		var previousNode *node
		var nextNode *node
		for currentNode != nil {
			fmt.Println(currentNode, " currentNode")

			nextNode = currentNode.next

			fmt.Println(nextNode, " nextNode")
			fmt.Println(previousNode, " previousNode")

			currentNode.next = previousNode

			fmt.Println(currentNode, " currentNode.next")

			previousNode = currentNode

			fmt.Println(previousNode, " previousNode")

			currentNode = nextNode

			fmt.Println(currentNode, " currentNode")

			fmt.Println()
		}
		l.head = previousNode
	}

	l.printList()
}

func (l *list) find(data int) {
	currentNode := l.head
	position := 0
	for currentNode != nil {
		position++
		if currentNode.data != data {
			currentNode = currentNode.next
			continue
		} else {
			fmt.Printf("%d is found at node %d\n", currentNode.data, position)
			break
		}
		fmt.Println("node not found")
	}
}

func (l *list) sort() {

	for i := 0; i < l.length; i++ {
		currentNode := l.head
		nextNode := l.head.next
		for j := 0; j < l.length-i-1; j++ {
			if currentNode.data > nextNode.data {
				temp := currentNode.data
				currentNode.data = currentNode.next.data
				currentNode.next.data = temp
			}
			currentNode = currentNode.next
			nextNode = nextNode.next
		}
	}
	l.printList()
}

func (l *list) insertAtPosition(newNode *node, pos int) {
	currentNode := l.head
	var previousNode *node
	var nextNode *node
	if pos > l.length || pos == 0 {
		panic(errors.New("invalid position"))
	} else {
		if pos == 1 {
			l.insertFirst(newNode)
			os.Exit(1)
		}
		for i := 1; i < pos; i++ {
			previousNode = currentNode
			currentNode = currentNode.next
		}
		nextNode = currentNode
		currentNode = newNode
		previousNode.next = currentNode
		currentNode.next = nextNode
		l.length++
	}
	l.printList()
}

func main() {
	list := &list{}
	n1 := node{
		data: 10,
	}
	n2 := node{
		data: 20,
	}
	n3 := node{
		data: 30,
	}
	list.append(&n1)
	list.append(&n2)
	list.append(&n3)
	list.reverseList()
	fmt.Println("length of the list = ", list.length)

	list.find(30)
	list.sort()
	list.delete(30)
	list.delete(10)
	n4 := node{
		data: 40,
	}
	list.insertFirst(&n4)
	n5 := node{
		data: 500,
	}
	list.insertFirst(&n5)
	list.append(&node{76, nil})
	n6 := node{
		data: 49,
	}
	list.insertAtPosition(&n6, 3)
}
