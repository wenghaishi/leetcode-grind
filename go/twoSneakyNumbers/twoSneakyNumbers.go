package main

import "fmt"

func twoSneakyNumbers(num []int) []int {
	var result []int
	frequency := make(map[int]struct{})

	for _, i := range num {
		_, exists := frequency[i]

		if exists {
			result = append(result, i)
		} else {
			frequency[i] = struct{}{}
		}
	}

	return result
}

func main(){
	test1 := []int{1,2,3,4,4,5,5}
	fmt.Println(twoSneakyNumbers(test1))

	test2 := []int{1,1,2,2,3,4,5 }
	fmt.Println(twoSneakyNumbers(test2))
}