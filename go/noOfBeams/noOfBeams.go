package main

import "fmt"

func noOfBeams(bank []string) int {
	
	var devices []int
	for _, i := range bank {
		allZero := true
		deviceCount := 0
		for j := 0; j < len(i); j++ {
			if i[j] == '1' {
				allZero = false
				deviceCount++
			}
		}
		if !allZero {
			devices = append(devices, deviceCount)
		}
	}

	solution := 0
	for k := 0; k < len(devices) - 1; k++ {
		solution += devices[k] * devices[k+1]
	}

	return solution
}

func main() {
	test1 := []string{"001", "000", "100"}
	fmt.Println(noOfBeams(test1))

	test2 := []string{"111", "111", "1", "001"}
	fmt.Println(noOfBeams(test2))
}