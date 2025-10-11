package main

import "fmt"

func containerWithMostWater (height []int) int {
	left, right := 0, len(height) - 1
	maxVol := 0

	for left < right {
		width := right - left
		vol := min(height[left], height[right]) * width

		if vol > maxVol {
			maxVol = vol
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxVol
}

func main() {
	arr := []int{1,2,3,4}
	fmt.Println(containerWithMostWater(arr))
}