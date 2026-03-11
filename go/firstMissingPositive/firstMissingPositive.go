package main

import (
	"math"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)

	if nums[0] > 1 {
		return 1
	}

	currentMax := 0

	for i := 0; i < len(nums) -1; i++ {
		if nums[i] <= 0 {
			continue
		} else if nums[i] + 1 == nums[i+1] {
			currentMax = 
		}
	}
	

}
