package main

import "fmt"

func main() {
	test1 := []int{10, 9, 1, 2, 3, 5, 7, 6, 4, 8}
	res := mergeSort(test1)
	fmt.Println(res)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	leftArr, rightArr := arr[:middle], arr[middle:]

	leftArr = mergeSort(leftArr)
	rightArr = mergeSort(rightArr)

	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	solution := make([]int, 0, len(leftArr)+len(rightArr))
	l, r := 0, 0

	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] < rightArr[r] {
			solution = append(solution, leftArr[l])
			l++
		} else {
			solution = append(solution, rightArr[r])
			r++
		}
	}

	solution = append(solution, leftArr[l:]...)
	solution = append(solution, rightArr[r:]...)

	return solution
}
