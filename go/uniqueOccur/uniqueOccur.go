package main

import "fmt"

func uniqueOccur(nums []int)bool {
	freq := make(map[int]int)

	for _, i := range nums {
		freq[i]++
	}

	seen := make(map[int]struct{})

	for _, value := range freq {
		_, exists := seen[value]

		if exists {
			return false
		} else {
			seen[value] = struct{}{}
		}
	}

	return true
}


func main() {
	test1 := []int{1,2,3}
	test2 := []int{1,2,2,3,3,3}

	fmt.Println(uniqueOccur(test1))

	fmt.Println(uniqueOccur(test2))

}