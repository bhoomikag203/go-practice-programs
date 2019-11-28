//implementation of selection sort
//time complexity is O(n^2)

package selectionSort

func selectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		minIndex := i

		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		list[minIndex], list[i] = list[i], list[minIndex]
	}
	return list
}

/*
func main() {
 	list := []int{64, 25, 12, 22, 11}
	fmt.Println(selectionSort(list))
}
*/
