// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

package main

import "fmt"

func jump(nums []int) int {
    result := 0
    left := 0
    right := 0

    for right < len(nums)-1 {
        farthest := 0

        for i := left; i <= right; i++ {
            farthest = max(farthest, i+nums[i])
        }

        left = right + 1
        right = farthest
        result++
    }

    return result
}

func main() {
    fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

