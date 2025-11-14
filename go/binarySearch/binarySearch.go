package main

import "fmt"

func bSearch(arr []int, target int) int {

	if len(arr) == 0 {
		return -1
	}

	bottom, top := 0, len(arr) - 1

	for top >= bottom {
		middle := bottom + top / 2

		if arr[middle] > target {
			top = middle
		} else if arr[middle] < target {
			bottom = middle
		} else if arr[middle] == target {
			fmt.Println("horray! you found it!")
			return middle
		}
	}

	return -1
}


func main() {
	test1 := []int{1,2,3,4,5,6,7}
	fmt.Println(bSearch(test1, 4))

	test2 := []int{1}
	fmt.Println(bSearch(test2, 1))
}
