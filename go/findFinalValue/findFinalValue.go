package main

import "fmt"

func findFinalValue(nums []int, original int) int {
	freq := make(map[int]struct{}, len(nums))

    for _, v := range nums {
        freq[v] = struct{}{}
    }

    for {
        if _, exists := freq[original]; exists {
            original *= 2
        } else {
            return original
        }
    }
}


func main() {
	fmt.Println(findFinalValue([]int{1,2,3,4,5}, 2))
}