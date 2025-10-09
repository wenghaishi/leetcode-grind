package main

import "fmt"

func canPlaceFlowers (arr []int, n int) bool {
	if n <= 0 {
		return true
	}

	if len(arr) == 1 {
		return arr[0] ==0
	}


	for i := range arr {
		if i == 0 {
			if arr[i] == 0 && arr[i+1] == 0 {
				n--
				arr[i] = 1
			}
		} else if i == len(arr) - 1 {
			if arr[i-1] == 0 && arr[i] == 0 {
				n--
				arr[i] = 1
			}
		} else if arr[i] == 0 && arr[i-1] == 0 && arr[i+1] == 0 {
			n--
			arr[i] = 1
		}
	}


	return n <= 0
}

func main () {
	arr := []int{0,0,0}
	fmt.Println(canPlaceFlowers(arr, 1))
}