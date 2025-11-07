package main

import "fmt"

func sellingStock(p []int) int {
	left, right := 0, 1
	var maxProfit int

	for right < len(p) {
		profit := p[right]- p[left]
		if profit < 0 {
			left = right
		} else {
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		right++
	}

	return maxProfit
}

func main () {
	test1 := []int{1,2,3,4,5}

	fmt.Println(sellingStock(test1))
}
