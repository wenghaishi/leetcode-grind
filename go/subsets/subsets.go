package main

import "fmt"

func subsets(arr []int) [][]int {
	solution := [][]int{{}}

	var helper func(index int, currentArr []int)
	helper = func(index int, currentArr []int) {
		for i := index; i
	}

	helper(0, arr)

	return solution
}

func main() {
	test := subsets([]int{1, 2, 3})
	fmt.Println(test)
}
1