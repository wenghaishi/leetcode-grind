package main

import "fmt"

func longestIncreasingSeq(nums []int) []int {
	maxLength := 0
	var left int
	var right int
	var maxSlice []int

	for i := 0; i < len(nums); i++ {
		left = i
		right = i

		for j := i; j < len(nums) - 1; j++ {
			if nums[j+1] == nums[j]+1 {
				fmt.Println("conseq")
				right++
				maxLength++
			} else {
				currentLength := right - left + 1

				fmt.Println("current l:", currentLength)

				if currentLength > maxLength {
					maxSlice = []int{}
					for k := i; k < right; k++ {
						maxSlice = append(maxSlice, nums[k])
					}
				}
			}
		}
	}

	return maxSlice
}



func main() {
	test1 := []int{1,2,3,6,7,1}

	fmt.Println(longestIncreasingSeq(test1))
}