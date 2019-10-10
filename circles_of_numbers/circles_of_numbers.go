package main

// import "fmt"

func CircleOfNumbers(numberOfParts, partNumber int) (oppositePart int) {
	oppositePart = (partNumber + numberOfParts/2) % numberOfParts
	return oppositePart
}

// func main() {
// 	fmt.Print("hello")
// }
