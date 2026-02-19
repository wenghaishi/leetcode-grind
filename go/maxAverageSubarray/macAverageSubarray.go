package main

func findMaxAverage(nums []int, k int) float64 {
		maxSum := 0
		current := k

		for current < len(nums) {
			sum := 0

			for i := current - k; i < k; i++ {
				sum+= nums[i]
			}

			if sum > maxSum {
				maxSum = sum
			}

			current++
		}

		return float64(maxSum) / float64(k)


}