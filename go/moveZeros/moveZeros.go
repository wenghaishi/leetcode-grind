package main

func moveZeros(nums []int) {
	pos := 0

	for _, i := range nums {
		if i != 0 {
			nums[pos] = i
			pos++
		}
	}

	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}