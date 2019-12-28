//implementation of insertion sort
//time complexity is O(n^2)

package insertion_sort

func insertion_sort(list []int) []int {
	for i := 0; i < len(list); i++ {
		key := list[i]
		j := i - 1

		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = key
	}
	return list
}

// func main() {
// 	list := []int{100, 64, 25, 12, 22, 11}
// 	fmt.Println(insertion_sort(list))
// }
