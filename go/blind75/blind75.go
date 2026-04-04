package main

import "fmt"

func main() {
	fmt.Println("Two sum result: ", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Best time to buy and sell stock: ", bestTimeToBuyAndSellStock([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Contains duplicate: ", containsDuplicate([]int{1, 2, 3, 1}))
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
	if len(nums) == 1 {
		return nums
	}

	solution := make([]int, 0)
	totalProduct := nums[0]
	numOfZeros := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			numOfZeros++
			continue
		}
		totalProduct *= nums[i]
	}

	if numOfZeros > 1 {
		for j := 0; j < len(nums); j++ {
			solution = append(solution, 0)
		}

		return solution
	}

	for k := 0; k < len(nums); k++ {
		if numOfZeros == 1 && nums[k] == 0 {
			solution = append(solution, totalProduct)
		} else {
			solution = append(solution, totalProduct/nums[k])
		}
	}

	return solution

}
