package main

import (
	"fmt"
	"slices"
)

func triangle(t [][]int) int {
	var sums [] int

	var helper func(row, rowPosition, currentSum int)

	helper = func(row, rowPosition, currentSum int) {

		// last row
		if row == len(t) -1 {
			sums = append(sums, currentSum + t[row][rowPosition])
		} else {
			newCurrentSum := currentSum + t[row][rowPosition]
			helper(row + 1, rowPosition, newCurrentSum)
			helper(row + 1, rowPosition + 1, newCurrentSum)
		}
	}
	helper(0, 0, 0)

	return slices.Max(sums)
}

func optimised (t [][]int) int {


	var helper func(rowSum []int, rowIndex int) int
	helper = func(rowSum []int, rowIndex int) int {

		// First row, pick the smallest path
		if rowIndex == 0 {
			return slices.Min([]int{rowSum[0] + t[0][0], rowSum[1] + t[0][0]})
		} else {
			var currentRowSum []int
			currentRow := t[rowIndex]
			for i := 0; i < len(currentRow); i++ {
				option1 := currentRow[i] + rowSum[i]
				option2 := currentRow[i] + rowSum[i+1]
				min := slices.Min([]int{option1, option2})

				currentRowSum = append(currentRowSum, min)
			}

			return helper(currentRowSum, rowIndex - 1)

		}
	}

	return helper(t[len(t) -1], len(t) -2)
}

func main() {
	fmt.Println(optimised([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}