//implementation of selection sort
//time complexity is O(n^2)

package selection_sort

func selectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		//assigns the first position as minimum element position
		minIndex := i

		//finds the minimum element in each iteration and updates the minimum(smaller) element's position
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		//once the smaller element position is found, then swap the smaller element
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
