package main

func removeDups(nums []int)int {
	current := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			current++
			nums[current] = nums[i]
		}
	}

	return current
}
