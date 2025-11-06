package main

import "fmt"

func sortColors(nums []int) {
	freq := make(map[int]int)

	for _, i := range nums {
		freq[i]++
	}

	currentIndex := 0
	for j := 0; j <= 2; j++ {
		numberValue := freq[j]

		for k := 0; k < numberValue; k++ {
			nums[currentIndex] = j
			currentIndex++
		}
	}
}

func main() {
	test1 := []int{2,1,0}
	sortColors(test1)
	fmt.Println(test1)
}