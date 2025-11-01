package main

import (
	"fmt"
	"math"
)

func findMaxAvg(nums []int, k int) float64 {
	maxAvg := math.Inf(-1)

	left, right := 0, k-1

	var initialSum int
	for i := 0; i < k; i++ {
		initialSum += nums[i]
	}

	for right < len(nums) {

        if left != 0 {
            initialSum -= nums[left - 1]
            initialSum += nums[right]
        }

		average := float64(initialSum) / float64(k)

		if average > maxAvg {
			maxAvg = average
		}

		left++
		right++
	}
	return maxAvg
}


func main() {
	test1 := []int{1,12,-5,-6,50,3}

	fmt.Println(findMaxAvg(test1, 4))
}
