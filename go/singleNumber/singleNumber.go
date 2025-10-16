package main

import "fmt"

func singleNumber(n []int) int {
	frequencies := make(map[int]int)

	for _, i := range n {
		_, exists := frequencies[i]

		if exists {
			delete(frequencies, i)
		} else {
			frequencies[i] = 1
		}
	}

	for key, _ := range frequencies {
		return key
	}


	return -1
}


func main() {
	test1 := []int{1,1,2,2,3}
	test2 := []int{2,1,3,1,2}

	fmt.Println(singleNumber(test1))
	fmt.Println(singleNumber(test2))
}
