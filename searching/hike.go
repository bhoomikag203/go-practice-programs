package main

import (
	"fmt"
)

func hike(steps string) (count int) {
	sealevel := 0
	count = 0
	s := []rune(steps)
	for _, char := range s {
		if char == 'U' {
			sealevel = sealevel + 1
			if sealevel == 0 {
				count++
			}
		} else if char == 'D' {
			sealevel = sealevel - 1
		}
	}
	return count
}

func main() {

	fmt.Println(hike("UDDDUDUU"))
}
