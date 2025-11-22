package main

func minOperations(nums []int) int {
	count := 0

	for _, i := range nums {
		operations := i % 3
		if operations == 2 {
			operations = 1
		}
		count += operations
	}

	return count
}