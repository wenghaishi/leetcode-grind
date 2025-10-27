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
				right++
			} else {
				currentLength := right - left + 1

				if currentLength > maxLength {
					maxSlice = []int{}
					maxLength = currentLength
					for k := i; k < right; k++ {
						maxSlice = append(maxSlice, nums[k])
					}
				} else {
					break
				}
			}
		}
	}

	return maxSlice
}

func optimised(nums []int)[]int{
	currentIndex := 0
	var maxLength int
	var currentLongestSlice [] int

	for  {
		if currentIndex >= len(nums) {
			break
		}
		right := currentIndex

		
		//check each right element
		for i := currentIndex; i < len(nums) - 1; i++ {
			if nums[i] + 1 == nums[i+1] {
				right++
			} else {
				currentLength := right - currentIndex +1

				if currentLength > maxLength {
					maxLength = currentLength
					currentLongestSlice = nums[currentIndex:right + 1]
				} else {
					currentIndex += currentLength
				}
			}
		}
	}

	return currentLongestSlice
}



func main() {
	test1 := []int{1,2,3,6,7,1}

	// fmt.Println(longestIncreasingSeq(test1))

	fmt.Println(optimised(test1))
}