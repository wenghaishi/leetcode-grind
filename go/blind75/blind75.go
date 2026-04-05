package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Two sum result: ", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Best time to buy and sell stock: ", bestTimeToBuyAndSellStock([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Contains duplicate: ", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println("Product except self: ", productExceptSelf([]int{1, 2, 3}))
	fmt.Println("Max subarray: ", maxSubArray([]int{1, 2, 3}))
	fmt.Println("Max product: ", maxProduct([]int{1, 2, 3}))
	fmt.Println("Min rotated sorted arr: ", findMin([]int{1, 2, 3}))

}

func twoSum(nums []int, target int) []int {
	freq := make(map[int]int)
	solution := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		value, exists := freq[nums[i]]

		if exists {
			solution = []int{value, i}
			break
		} else {
			freq[target-nums[i]] = i
		}
	}

	return solution
}

func bestTimeToBuyAndSellStock(prices []int) int {
	l, r := 0, 1
	maxProfit := 0

	for r < len(prices) {

		profit := prices[r] - prices[l]
		if profit > maxProfit {
			maxProfit = profit
		}

		if prices[r] < prices[l] {
			l = r
		}

		r++
	}
	return maxProfit
}

func containsDuplicate(nums []int) bool {
	freq := make(map[int]struct{})

	for _, v := range nums {
		_, exists := freq[v]

		if exists {
			return true
		} else {
			freq[v] = struct{}{}
		}
	}

	return false
}

func productExceptSelf(nums []int) []int {
	solution := make([]int, len(nums))
	prefix, postfix := 1, 1

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			solution[i] = prefix
			continue
		}

		prefix *= nums[i-1]
		solution[i] = prefix
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			continue
		}

		postfix *= nums[j+1]
		solution[j] *= postfix
	}

	return solution
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := math.MinInt
	currentSum := 0
	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		if currentSum > maxSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}

func maxProduct(nums []int) int {
	res := math.MinInt
	currMax, currMin := 1, 1

	for _, v := range nums {
		tmp := currMax * v
		currMax = max(v, v*currMax, v*currMin)
		currMin = min(v, tmp, v*currMin)
		res = max(res, currMax)
	}

	return res
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := math.MaxInt
	for left <= right {
		if nums[left] < res {
			res = nums[left]
		}

		if nums[right] < res {
			res = nums[right]
		}
		left++
		right--
	}

	return res
}
