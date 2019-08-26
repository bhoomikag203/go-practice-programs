//to calculate (row with maximum sum + column with maximum sum)
package main

import (
	"fmt"
)

func Sum(matrix [][]int) int {
	// to calculate row with maximum sum
	max_row_sum := 0
	for i := 0; i < len(matrix); i++ {
		row_sum := 0
		for j := 0; j < len(matrix[i]); j++ {
			row_sum += matrix[i][j]
		}
		if row_sum > max_row_sum {
			max_row_sum = row_sum
		}

	}

	//to calculate column with maximum sum
	max_col_sum := 0
	for i := 0; i < len(matrix[1]); i++ {
		col_sum := 0
		for j := 0; j < len(matrix); j++ {
			col_sum += matrix[j][i]
		}
		if col_sum > max_col_sum {
			max_col_sum = col_sum
		}

	}
	fmt.Println("max row sum = ", max_row_sum)
	fmt.Println("max col sum = ", max_col_sum)

	sum := max_row_sum + max_col_sum
	return sum
}

func main() {
	matrix := [][]int{{1, 8, 6}, {4, 1, 5}}
	fmt.Println("Total = ", Sum(matrix))
}
