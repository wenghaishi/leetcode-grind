package main

import (
	"fmt"
	"math"
)

func increasingTriplets(nums []int) bool{
	smallest, secondSmallest := math.MaxInt, math.MaxInt

	for _, i := range nums {
		if i <= smallest {
			smallest = i
		} else if i <= secondSmallest {
			secondSmallest = i
		}  else {
			return true
		}
	}

	return false
}

func main () {
	test1 := []int{1,2,3}
	test2 := []int {3,2,1}
	test3 := []int{0,7,2,1,3}
	fmt.Println(increasingTriplets(test1))
	fmt.Println(increasingTriplets(test2))
	fmt.Println(increasingTriplets(test3))
}
