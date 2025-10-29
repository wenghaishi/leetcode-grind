package main

import "fmt"


func countElementsGreaterThanPrevAvg(arr []int) int {
	var currentSum int
	count := 0
	for i := 0; i < len(arr); i++ {
		if i != 0 {
			avg := float64(currentSum) / float64((i))

			if float64(arr[i]) > avg {
				count++
			}

			currentSum += arr[i]

		} else {
			currentSum += arr[i]
		}

	}

	return count
}

func main() {
	test1 := []int{1,2,3}
	fmt.Println(countElementsGreaterThanPrevAvg(test1))

	test2 := []int{3,10,1,1,1,1,1}
	fmt.Println(countElementsGreaterThanPrevAvg(test2))

}